package driver

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/Gujarats/nearest/util"

	"github.com/Gujarats/nearest/model/city/interface"
	driverInterface "github.com/Gujarats/nearest/model/driver/interface"

	driverModel "github.com/Gujarats/nearest/model/driver"

	"github.com/Gujarats/nearest/model/global"
)

// create logger to print error in the console
var logger *log.Logger

func init() {
	logger = log.New(os.Stderr,
		"Controller Driver :: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

// find specific driver with their ID or name.
// if the desired data didn't exist then insert new data
func UpdateDriver(m *sync.Mutex, driver driverInterface.DriverInterfacce, cityInterface cityInterface.CityInterfacce) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//start time for lenght of the process
		startTimer := time.Now()

		w.Header().Set("Access-Control-Allow-Methods", "POST")

		id := r.FormValue("id")
		name := r.FormValue("name")
		lat := r.FormValue("latitude")
		lon := r.FormValue("longitude")
		status := r.FormValue("status")
		city := r.FormValue("city")

		isAllExist := util.CheckValue(id, name, lat, lon, status, city)
		if !isAllExist {
			logger.Println("Required Params Empty")

			//return Bad response
			w.WriteHeader(http.StatusBadRequest)
			global.SetResponse(w, "Failed", "Required Params Empty")
			return
		}

		// check id for validation format
		ok := bson.IsObjectIdHex(id)
		if !ok {
			//return Bad response
			w.WriteHeader(http.StatusBadRequest)
			global.SetResponse(w, "Failed", "Invalid Id format")
			return
		}

		// convert string to bool
		statusBool, err := strconv.ParseBool(status)
		if err != nil {
			//return Bad response
			w.WriteHeader(http.StatusBadRequest)
			global.SetResponse(w, "Failed", "Parse Boolean Error")
			return
		}

		// convert string to float64
		convertedFloat, err := util.ConvertToFloat64(lat, lon)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			global.SetResponse(w, "Failed", "Failed to conver float value")
			return
		}

		latFloat := convertedFloat[0]
		lonFloat := convertedFloat[1]

		// check driver last location which district they were before
		lastDistrict := driver.GetLastDistrict(id)

		// checks drivers location which district they are now
		// NOTE : getting the nearest district must not null or fail. so we need to repeat the function if we got null.
		// but there is one approach solutions we give more distance value so that we can find the district event it is far.
		district, err := cityInterface.GetNearestDistrict(city, latFloat, lonFloat, 500)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			global.SetResponse(w, "Failed", "Failed to get nearest district")
			return

		}

		// checking if we got the current district data
		if district.Name == "" {
			w.WriteHeader(http.StatusOK)
			global.SetResponse(w, "Success", "No nearest district found!")
			return

		}

		// check if the location is the same or not, if not then remove the data from the lastdistrict.
		// format current district to meet the last district format
		if lastDistrict != "" {
			// lastDistrict is not empty from redis check if it is the same as current location
			if district.Name+"_district_"+district.Id.Hex() != lastDistrict {
				// remove the driver data in the last district
				// lastDistrict must formatted like collectionKey for district collections
				driver.Remove(id, lastDistrict)
			}

		}

		driverData := driverModel.DriverData{Id: bson.ObjectIdHex(id), Name: name, Status: statusBool, Location: driverModel.GeoJson{Type: "Point", Coordinates: []float64{lonFloat, latFloat}}}
		err = driver.Update(district.Name, district.Id.Hex(), driverData)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			global.SetResponse(w, "Failed", "Failed to update the driver")
			return
		}

		//return succes response
		elpasedTime := time.Since(startTimer).Seconds()
		w.WriteHeader(http.StatusOK)
		global.SetResponseTime(w, "Succes", "Driver Updated", elpasedTime)
		return
	})

}

type request func(w http.ResponseWriter, r *http.Request)

func FindDriver(m *sync.Mutex, driver driverInterface.DriverInterfacce, cityInterface cityInterface.CityInterfacce) request {

	return func(w http.ResponseWriter, r *http.Request) {

		m.Lock()
		defer m.Unlock()

		//start time for lenght of the process
		startTimer := time.Now()

		lat := r.FormValue("latitude")
		lon := r.FormValue("longitude")
		city := r.FormValue("city")
		distance := r.FormValue("distance")

		//checking empty value
		checkValue := util.CheckValue(lat, lon, city, distance)
		if !checkValue {
			logger.Println("Required Params Empty")

			//return Bad response
			w.WriteHeader(http.StatusBadRequest)
			global.SetResponse(w, "Failed", "Required Params Empty")
			return
		}

		floatNumbers, err := util.ConvertToFloat64(lat, lon)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			global.SetResponse(w, "Failed", "Failed to convert float value")
			return
		}
		latFloat := floatNumbers[0]
		lonFloat := floatNumbers[1]

		intNumbers, err := util.ConvertToInt64(distance)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			global.SetResponse(w, "Failed", "Failed to convert integer value")
			return
		}
		distanceInt := intNumbers[0]

		// determined which quadran the input locations

		// NOTE : send the driver which locations and its distance from input location
		// 1.  getting all centers location map[int][4]locations
		// 2. determined which quadran is it for inputLocation
		// 3. get nearest marked locations from that quadran
		// 4. get driver from the collections of nearest locations(this is must be array)
		// 5. send the driver
		// 6. if the driver is not found then go to next nearest locations. currentIndex +1. until the last index.
		// 7. if not found find the driver on the next quadran (currentquadran +1) until reach 4
		// 8. if not found find the driver on the next level. repeat this until you found the driver

		// get all district from redis and calculate it
		// calculate nearest location district with given location and city from mongodb
		district, err := cityInterface.GetNearestDistrict(city, latFloat, lonFloat, distanceInt)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			global.SetResponse(w, "Failed", "Failed to get nearest district")
			return
		}

		// checking district result from mongodb
		if district.Name == "" {
			w.WriteHeader(http.StatusInternalServerError)
			global.SetResponse(w, "Failed", "No district found")
			return
		}

		//response variable for getting the drivers
		var driverResponse driverModel.DriverData

		// checks drivers int the district from the redis
		drivers := driver.DriversRedis(district.Name, district.Id.Hex())
		if len(drivers) > 0 {
			// get the first index drvier from redis and save it again to redis
			driverResponse = drivers[0]

			driver.SaveLastDistrict(drivers[0].Id.Hex(), district.Name, district.Id.Hex())

			// update the driver's status to unavailable in mongodb
			// Latitude is 1 in the index and Longitude is 0. Rules from mongodb
			drivers[0].Status = false
			err := driver.Update(district.Name, district.Id.Hex(), drivers[0])
			if err != nil {

			}

			// update redis data by removing the first index
			drivers = drivers[1:]
			// save the drivers to redis replacing previous data
			driver.SaveDriversRedis(drivers, district.Name, district.Id.Hex())

		} else {
			// get drivers from mongodb
			drivers = driver.GetAvailableDriver(district.Name, district.Id.Hex())
			if len(drivers) > 0 {
				driverResponse = drivers[0]

				drivers[0].Status = false
				err := driver.Update(district.Name, district.Id.Hex(), drivers[0])
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					global.SetResponse(w, "Failed", "Failed to update ther driver")
					return
				}

				// update redis data by removing the first index
				drivers = drivers[1:]
				// save the drivers to redis replacing previous data
				driver.SaveDriversRedis(drivers, district.Name, district.Id.Hex())
			} else {
				// we could not find any data in redis and mongo
				w.WriteHeader(http.StatusOK)
				global.SetResponse(w, "Success", "We couldn't find any driver")
				return
			}
		}

		//return succes response
		w.WriteHeader(http.StatusOK)
		elapsedTime := time.Since(startTimer).Seconds()
		response := global.Response{Status: "Success", Message: "Data found", Latency: elapsedTime, Data: driverResponse}
		json.NewEncoder(w).Encode(response)
		return

	}
}

func InsertDriver(driver driverInterface.DriverInterfacce) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//start time for lenght of the process
		startTimer := time.Now()
		w.Header().Set("Access-Control-Allow-Methods", "POST")

		// getting the parameters
		name := r.FormValue("name")
		lat := r.FormValue("latitude")
		lon := r.FormValue("longitude")
		status := r.FormValue("status")

		isAllExist := util.CheckValue(name, lat, lon, status)
		if !isAllExist {
			logger.Println("Required Params Empty")

			//return Bad response
			w.WriteHeader(http.StatusBadRequest)
			global.SetResponse(w, "Failed", "Required Params Empty")
			return
		}

		// convert string to bool
		statusBool, err := strconv.ParseBool(status)
		if err != nil {
			//return Bad response
			logger.Println("Failed to Parse Boolean")
			w.WriteHeader(http.StatusBadRequest)
			global.SetResponse(w, "Failed", "Parse Boolean Erro")
			return
		}

		// convert string to float64
		convertedFloat, err := util.ConvertToFloat64(lat, lon)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			global.SetResponse(w, "Failed", "Failed to convert float value")
			return
		}
		latFloat := convertedFloat[0]
		lonFloat := convertedFloat[1]

		// insert driver t
		// TODO :  error handler here what happens if process of insertion data fails
		driver.Insert(name, name, latFloat, lonFloat, statusBool)

		//return succes response
		w.WriteHeader(http.StatusOK)
		elapsedTime := time.Since(startTimer).Seconds()
		response := global.Response{Status: "Success", Message: "Data Inserted", Latency: elapsedTime}
		json.NewEncoder(w).Encode(response)
		return
	})
}
