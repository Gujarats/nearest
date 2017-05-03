package main

// NOTE :  This file is used for seed the database. Here I'm creating 2500 district or marked location based on one city. from given latitude and longitude.
// in every district I created 1000 drivers so that user in near district can request a driver.

import (
	"fmt"
	"log"

	// fake data library
	"github.com/Gujarats/GenerateLocation"
	"github.com/icrowley/fake"

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
	insertDummyMarkLocation(cityName, city)

	// inserting dummy driver
	insertDummyDriver(cityName, city, driverData)
}

// insert dummy location from latitude and longitude.
func insertDummyMarkLocation(cityName string, city *cityModel.City) {
	var cities []cityModel.City
	// some location in Bandung that will be the top left corner base location.
	lat := -6.8647721
	lon := 107.553501
	var locations []location.Location

	// geneerate location with distance 1 km in every point and limit lenght 50 km.
	// so it will be (50/1)^2 = 2500 district
	locations = location.GenerateLocation(lat, lon, 0.5, 15.0)

	err := city.CreateIndex(cityName)
	if err != nil {
		log.Panic(err)
	}
	for index, resultLocation := range locations {
		//err := city.InsertDistrict(cityName, index, resultLocation.Lat, resultLocation.Lon)
		//if err != nil {
		//	log.Panic(err)
		//}
		cityData := cityModel.City{
			Name:     cityName,
			District: index,
			Location: cityModel.GeoJson{Type: "Point", Coordinates: []float64{resultLocation.Lon, resultLocation.Lat}},
		}

		cities = append(cities, cityData)
	}

	err = city.InsertDistrictBulk(cityName, cities)
	if err != nil {
		log.Panic(err)
	}
}

// insert 2.500.000 drivers. 1000 drivers in every district.
// passed driver struct to save the data to database.
func insertDummyDriver(cityName string, city *cityModel.City, driverData *driverModel.DriverData) {

	var drivers []driverModel.DriverData

	// getting all district from a city
	districts, err := city.AllDistrict(cityName)
	if err != nil {
		log.Panic(err)
	}

	// create one
	for _, district := range districts {
		// generate 1000 drivers
		fmt.Printf("Generating drivers in distrtic = %+v\n", district)
		dummyDrivers := GenereateDriver(district.Location.Coordinates[1], district.Location.Coordinates[0], 1000)

		//create collectionName for using the format: cityName_district_DistrictId
		districtId := district.Id.Hex()
		collectionsName := cityName + "_district_" + districtId

		for _, driver := range dummyDrivers {

			driverStruct := driverModel.DriverData{
				Name:     driver.Name,
				Status:   driver.Status,
				Location: driverModel.GeoJson{Type: "Point", Coordinates: []float64{driver.Lon, driver.Lat}},
			}

			// append driverStruct to drivers to bulk data
			drivers = append(drivers, driverStruct)

			// create index driver
			//driverData.Insert(collectionsName, driver.Name, driver.Lat, driver.Lon, driver.Status)
		}

		err := driverData.CreateIndex(collectionsName)
		if err != nil {
			log.Panic(err)
		}

		driverData.InsertBulk(collectionsName, drivers)
		drivers = nil
	}
}

// this struct is used for GenereteDriver
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
	location.SetupLocation(lat, lon)

	// get 50 % of the sum data
	smallPercentage := (50.0 / 100.0) * float64(sum)
	percentData := int(smallPercentage)

	// random lat lon based on seconds
	for i := 0; i <= sum; i++ {
		if sum-i <= percentData {
			// generate lat and lon using minute. from specific number 1-3
			lat, lon := location.RandomLatLongMinute(4)
			dummyDriver := Driver{
				Name:   fake.FullName(),
				Lat:    lat,
				Lon:    lon,
				Status: false,
			}
			drivers = append(drivers, dummyDriver)
		} else {
			// generate lat and lon using seconds. from specific number 1-6
			lat, lon := location.RandomLatLong(7)
			dummyDriver := Driver{
				Name:   fake.FullName(),
				Lat:    lat,
				Lon:    lon,
				Status: true,
			}
			drivers = append(drivers, dummyDriver)
		}

	}

	return drivers
}
