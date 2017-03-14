package driver

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/Gujarats/API-Golang/util"
	"github.com/Gujarats/API-Golang/util/logger"

	driverModel "github.com/Gujarats/API-Golang/model/driver"
	"github.com/Gujarats/API-Golang/model/driver/interface"
	"github.com/Gujarats/API-Golang/model/global"
)

// find specific driver with their ID or name.
// if the desired data didn't exist then insert new data
func UpdateDriver(driver driverInterface.DriverInterfacce) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//start time for lenght of the process
		startTimer := time.Now()

		w.Header().Set("Access-Control-Allow-Methods", "POST")

		id := r.FormValue("id")
		name := r.FormValue("name")
		lat := r.FormValue("latitude")
		lon := r.FormValue("longitude")
		status := r.FormValue("status")

		isAllExist := util.CheckValue(id, name, lat, lon, status)
		if !isAllExist {
			logger.PrintLog("Required Params Empty")

			//return Bad response
			w.WriteHeader(http.StatusBadRequest)
			global.SetResponse(w, "Failed", "Required Params Empty")
			return
		}

		// convert string to bool
		statusBool, err := strconv.ParseBool(status)
		if err != nil {
			//return Bad response
			logger.PrintLog("Failed to Parse Boolean")
			w.WriteHeader(http.StatusBadRequest)
			global.SetResponse(w, "Failed", "Parse Boolean Erro")
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

		driverData := driverModel.DriverData{Id: bson.ObjectId(id), Name: name, Status: statusBool, Location: driverModel.GeoJson{Coordinates: []float64{lonFloat, latFloat}}}
		driver.Update(driverData)

		//return succes response
		elpasedTime := time.Since(startTimer).Seconds()
		w.WriteHeader(http.StatusOK)
		global.SetResponseTime(w, "Succes", "Driver Inserted", elpasedTime)
		return
	})

}

func FindDriver(driver driverInterface.DriverInterfacce) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//start time for lenght of the process
		startTimer := time.Now()
		w.Header().Set("Access-Control-Allow-Methods", "GET")

		lat := r.FormValue("latitude")
		lon := r.FormValue("longitude")
		city := r.FormValue("city")
		distance := r.FormValue("distance")

		//checking empty value
		checkValue := util.CheckValue(lat, lon, city, distance)
		if !checkValue {
			logger.PrintLog("Required Params Empty")

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

		//driver data for response
		var driverResponse driverModel.DriverData
		// get driver from redis
		currentDate := time.Now().Local().Format("01-02-2016")
		key := city + ";" + currentDate
		cityRedis, drivers := driver.DriversRedis(key)
		if cityRedis.Name != "" && len(drivers) > 0 {
			// calculate distance with the saved location in redis
			distance := util.Distance(cityRedis.Lat, cityRedis.Lon, latFloat, lonFloat)
			if distance <= 100.0 {
				// get the first index drvier from redis and save it again to redis
				driverResponse = drivers[0]
				// update the driver's status to unavailable in mongodb
				// Latitude is 1 in the index and Longitude is 0. Rules from mongodb
				driver.Update(drivers[0])
				// update redis data by removing the first index
				drivers = drivers[1:]
				// save the drivers to redis replacing previous data
				driver.SaveDriversRedis(drivers, cityRedis)

			}

		} else {
			// get the nearest drivers from mongod and get the first index
			drivers := driver.GetNearLocation(distanceInt, latFloat, lonFloat)
			if len(drivers) > 0 {
				driverResponse = drivers[0]

				// remove the first index to update to redis
				drivers = drivers[1:]

				// create city location based on the incoming request
				cityRedis := driverModel.City{Name: city, Lat: latFloat, Lon: lonFloat}
				driver.SaveDriversRedis(drivers, cityRedis)

				// update the driver's status to unavailable in mongodb
				// Latitude is 1 in the index and Longitude is 0. Rules from mongodb
				driver.Update(driverResponse)

			} else {
				// we could not find any data in redis and mongo
				w.WriteHeader(http.StatusOK)
				global.SetResponse(w, "Success", "We couldn't find any data")
				return
			}
		}

		//return succes response
		w.WriteHeader(http.StatusOK)
		elapsedTime := time.Since(startTimer).Seconds()
		response := global.Response{Status: "Success", Message: "Data Found", Latency: elapsedTime, Data: driverResponse}
		json.NewEncoder(w).Encode(response)
		return

	})
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
			logger.PrintLog("Required Params Empty")

			//return Bad response
			w.WriteHeader(http.StatusBadRequest)
			global.SetResponse(w, "Failed", "Required Params Empty")
			return
		}

		// convert string to bool
		statusBool, err := strconv.ParseBool(status)
		if err != nil {
			//return Bad response
			logger.PrintLog("Failed to Parse Boolean")
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

		// insert driver
		driver.Insert(name, latFloat, lonFloat, statusBool)

		//return succes response
		w.WriteHeader(http.StatusOK)
		elapsedTime := time.Since(startTimer).Seconds()
		response := global.Response{Status: "Success", Message: "Data Inserted", Latency: elapsedTime}
		json.NewEncoder(w).Encode(response)
		return
	})
}
