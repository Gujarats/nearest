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
	driverData.GetConn(mongoConn)

	// inserting dummy driver
	insertDummyDriver(driverData)

	//seed redis
	startTime := time.Now()
	seedDriverRedis(redisConn)
	GetDriverRedis(redisConn)
	fmt.Printf("takes time = %.6f\n", time.Since(startTime).Seconds())

}

// insert database 50.000 rows
// passed driver struct to save the data to database.
func insertDummyDriver(driverData *driverModel.DriverData) {

	dummyDrivers := GenereateDriver(50000)
	for _, driver := range dummyDrivers {
		driverData.Insert(driver.Name, driver.Lat, driver.Lon, driver.Status)
	}

}

func seedDriverRedis(redisConn *redis.Client) {
	// create 20 driver
	drivers := GenereateDriver(20)
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

	//fmt.Printf("city = %+v\n", city)
	//fmt.Printf("drivers = %+v\n", drivers)
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

func GenereateDriver(sum int) []Driver {
	var drivers []Driver
	location.SetupLocation(48.8588377, 2.2775176)

	// get 30 % of the sum data
	smallPercentage := (30.0 / 100.0) * float64(sum)
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
				Status: true,
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
