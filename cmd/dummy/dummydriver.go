package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Gujarats/API-Golang/database"
	"github.com/Gujarats/GenerateLocation"
	"github.com/icrowley/fake"
	mgo "gopkg.in/mgo.v2"
	redis "gopkg.in/redis.v5"

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

	//seed redis
	startTime := time.Now()
	seedDriverRedis(redisConn)
	GetDriverRedis(redisConn)
	fmt.Printf("takes time = %.6f\n", time.Since(startTime).Seconds())

}

// insert dummy location from latitude and longitude.
func insertDummyMarkLocation(cityName string, city *cityModel.City) {
	// some location in Bandung
	lat := -6.8647721
	lon := 107.553501
	var locations []location.Location

	// geneerate location with distance 1 km in every point and limit lenght 50 km.
	// so it will be (50/1)^2 = 2500 district
	locations = location.GenerateLocation(lat, lon, 1, 50)

	err := city.CreateIndex(cityName)
	if err != nil {
		log.Panic(err)
	}
	for index, resultLocation := range locations {
		err := city.InsertDistrict(cityName, index, resultLocation.Lat, resultLocation.Lon)
		if err != nil {
			log.Panic(err)
		}
	}

}

// insert 50.000 drivers. 100 drivers in every district.
// passed driver struct to save the data to database.
func insertDummyDriver(cityName string, city *cityModel.City, driverData *driverModel.DriverData) {

	// getting all district from a city
	districts, err := city.AllDistrict(cityName)
	if err != nil {
		log.Panic(err)
	}

	// create one
	for _, district := range districts {
		// generate 100 drivers
		fmt.Printf("Generating drivers in distrtic = %+v\n", district)
		dummyDrivers := GenereateDriver(district.Location.Coordinates[1], district.Location.Coordinates[0], 1000)

		//create collectionName for using the format: cityName_district_DistrictId
		districtId := district.Id.Hex()
		collectionsName := cityName + "_district_" + districtId

		for _, driver := range dummyDrivers {

			// print the data

			// create index driver
			err := driverData.CreateIndex(collectionsName)
			if err != nil {
				log.Panic(err)
			}
			driverData.Insert(collectionsName, driver.Name, driver.Lat, driver.Lon, driver.Status)
		}
	}
}

func seedDriverRedis(redisConn *redis.Client) {
	// create 20 driver
	drivers := GenereateDriver(48.8588377, 2.2775176, 20)
	city := City{
		Name: "Paris",
		Lat:  48.8588377,
		Lon:  2.2775176,
	}

	byteDrivers, _ := json.Marshal(drivers)
	byteCity, _ := json.Marshal(city)
	data := string(byteCity) + "++" + string(byteDrivers)

	redisConn.Set("someKey", data, 0)
}

func GetDriverRedis(redisConn *redis.Client) {
	var city City
	var drivers []Driver

	dataString, err := redisConn.Get("someKey").Result()
	if err != nil {
		log.Panic(err)
	}

	// split data with ++
	dataSplit := strings.Split(dataString, "++")
	byteCity := []byte(dataSplit[0])
	byteDrivers := []byte(dataSplit[1])

	err = json.Unmarshal(byteCity, &city)
	err = json.Unmarshal(byteDrivers, &drivers)
}

type Driver struct {
	Name   string  `json:"name"`
	Lat    float64 `json:"lat"`
	Lon    float64 `json:"lon"`
	Status bool    `json:"status"`
}

type City struct {
	Name string  `json:"name"`
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
}

func GenereateDriver(lat, lon float64, sum int) []Driver {
	var drivers []Driver
	location.SetupLocation(lat, lon)

	// get 30 % of the sum data
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
