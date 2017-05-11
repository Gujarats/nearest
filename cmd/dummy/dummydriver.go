package main

// NOTE :  This file is used for seed the database. Here I'm creating 2500 district or marked location based on one city. from given latitude and longitude.
// in every district I created 1000 drivers so that user in near district can request a driver.

import (
	"fmt"
	"log"
	"strconv"
	"sync"

	// fake data library

	location "gopkg.in/gujarats/GenerateLocation.v1"

	mgo "gopkg.in/mgo.v2"
	redis "gopkg.in/redis.v5"

	"github.com/Gujarats/API-Golang/database"

	cityModel "github.com/Gujarats/API-Golang/model/city"
	driverModel "github.com/Gujarats/API-Golang/model/driver"
)

func main() {
	//getting list of all the connection.
	listConnection := database.SystemConnection()

	// get mongodb connection
	mongoConn := listConnection["mongodb"].(*mgo.Session)
	redisConn := listConnection["redis"].(*redis.Client)

	// init driver model
	driverData := &driverModel.DriverData{}
	driverData.GetConn(mongoConn, redisConn)

	// init city model
	cityName := "Bandung"
	city := &cityModel.City{}
	city.GetConn(mongoConn, redisConn)

	// inserting city district
	fmt.Println("Please wait generating locations")
	insertDummyMarkLocation(cityName, city)

	// inserting dummy driver
	fmt.Println("Please wait generating drivers")
	insertDummyDriver(mongoConn, cityName, city, driverData)
}

// insert dummy location from latitude and longitude.
func insertDummyMarkLocation(cityName string, city *cityModel.City) {
	var cities []cityModel.City
	// some location in Bandung that will be the top left corner base location.
	lat := -6.8647721
	lon := 107.553501

	loc := location.New(lat, lon)
	var locations []location.Location

	// geneerate location with distance 1 km in every point and limit lenght 50 km.
	// so it will be (15.0/0.5)^2 = 900 district
	locations = loc.GenerateLocation(0.5, 15.0)
	mapCenterLocations, err := loc.GetCenterQuadranLocations(0.5, 15.0, 3)
	if err != nil {
		log.Fatal(err)
	}

	err = city.CreateIndex(cityName)
	if err != nil {
		log.Panic(err)
	}

	for index, resultLocation := range locations {
		cityData := cityModel.City{
			Name:     cityName,
			District: index,
			Location: cityModel.GeoJson{Type: "Point", Coordinates: []float64{resultLocation.Lon, resultLocation.Lat}},
		}

		cities = append(cities, cityData)
	}

	// insertt all district to one collections
	datas := make([]interface{}, len(cities))
	for index, city := range cities {
		datas[index] = city
	}

	err = city.InsertLocationsBulk(cityName, datas)
	if err != nil {
		log.Panic(err)
	}

	// insert district with its quadran position
	insertLocationToItsQuadran(city, mapCenterLocations, locations)
}

// this function will insert the input locations to its quadran.
// NOTE : every quadran will have a duplicate or the same locations.
// because I split them using to 4 pieces like I,II,III,IV in quadran.
func insertLocationToItsQuadran(city *cityModel.City, mapCenterLocations map[int][4]location.CenterLocation, locations []location.Location) {
	// iterate map that holds key as level and centerLocations as value
	for level, centerLocations := range mapCenterLocations {
		// create map that will hold all location to their quadran level
		mapLocations := make(map[string][]interface{})

		// center locations from the map.
		for _, centerLocation := range centerLocations {
			baseLocation := location.Location{
				Lat: centerLocation.MarkedLocation.Lat,
				Lon: centerLocation.MarkedLocation.Lon,
			}

			// iterate all locations which is input location to inserted to its quadran level.
			for _, inputLocation := range locations {
				quadran, err := location.GetQuadranPosition(baseLocation, inputLocation)
				if err != nil {
					log.Fatal(err)
				}

				// store inputLocations to its quadran and level collections on map
				storeLocation := struct{ Location cityModel.GeoJson }{
					Location: cityModel.GeoJson{Type: "Point", Coordinates: []float64{inputLocation.Lon, inputLocation.Lat}},
				}
				collectionName := "level_" + strconv.Itoa(level) + quadran
				mapLocations[collectionName] = append(mapLocations[collectionName], storeLocation)
			}

		}

		// after all locations has stored to the map, we store them into mongodb
		var wg sync.WaitGroup
		for collectionName, datas := range mapLocations {
			wg.Add(1)
			go func(city *cityModel.City, collectionName string, datas []interface{}) {
				err := city.CreateIndex(collectionName)
				if err != nil {
					log.Panic(err)
				}

				err = city.InsertLocationsBulk(collectionName, datas)
				if err != nil {
					log.Fatal(err)
				}
				wg.Done()
			}(city, collectionName, datas)

		}
		wg.Wait()
	}
}

// insert 2.500.000 drivers. 1000 drivers in every district.
// passed driver struct to save the data to database.
func insertDummyDriver(mongoConn *mgo.Session, cityName string, city *cityModel.City, driverData *driverModel.DriverData) {
	// getting all district from a city
	districts, err := city.AllDistrict(cityName)
	if err != nil {
		log.Panic(err)
	}

	var wg sync.WaitGroup
	fmt.Println("length dristricts = ", len(districts))
	for _, district := range districts {
		//create collectionName for using the format: cityName_district_DistrictId
		districtId := district.Id.Hex()
		collectionName := cityName + "_district_" + districtId

		// for storing generated drivers
		var drivers []interface{}

		// generate 1000 drivers
		dummyDrivers := GenereateDriver(district.Location.Coordinates[1], district.Location.Coordinates[0], 1000)
		for _, driver := range dummyDrivers {

			driverStruct := driverModel.DriverData{
				Name:     driver.Name,
				Status:   driver.Status,
				Location: driverModel.GeoJson{Type: "Point", Coordinates: []float64{driver.Lon, driver.Lat}},
			}

			// append driverStruct to drivers to bulk data
			drivers = append(drivers, driverStruct)

		}

		wg.Add(1)
		go func(collectionName string, driverData *driverModel.DriverData) {
			defer wg.Done()
			err = driverData.InsertBulk(collectionName, drivers)
			if err != nil {
				log.Panic(err)
			}

			err := driverData.CreateIndex(collectionName)
			if err != nil {
				log.Panic(err)
			}

			drivers = nil

		}(collectionName, driverData)
	}

	wg.Wait()

}

// this struct is used for Generete Driver
type Driver struct {
	Name   string  `json:"name"`
	Lat    float64 `json:"lat"`
	Lon    float64 `json:"lon"`
	Status bool    `json:"status"`
}

// Generate dummy drivers this will return []Driver with given sum.
// with new location from latitude and longitude given.
func GenereateDriver(lat, lon float64, sum int) []Driver {
	var drivers []Driver
	loc := location.New(lat, lon)

	// get 50 % of the sum data
	smallPercentage := (50.0 / 100.0) * float64(sum)
	percentData := int(smallPercentage)

	// random lat lon based on seconds
	for i := 0; i <= sum; i++ {
		if sum-i <= percentData {
			// generate lat and lon using minute. from specific number 1-3
			lat, lon := loc.RandomLatLongMinute(4)
			dummyDriver := Driver{
				Name:   "DummyFalse",
				Lat:    lat,
				Lon:    lon,
				Status: false,
			}
			drivers = append(drivers, dummyDriver)
		} else {
			// generate lat and lon using seconds. from specific number 1-6
			lat, lon := loc.RandomLatLongSeconds(7)
			dummyDriver := Driver{
				Name:   "DummyTrue",
				Lat:    lat,
				Lon:    lon,
				Status: true,
			}
			drivers = append(drivers, dummyDriver)
		}

	}

	return drivers
}
