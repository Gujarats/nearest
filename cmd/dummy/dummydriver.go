package main

import (
	"github.com/Gujarats/API-Golang/database"
	"github.com/Gujarats/GenerateLocation"
	"github.com/icrowley/fake"
	mgo "gopkg.in/mgo.v2"

	driverModel "github.com/Gujarats/API-Golang/model/driver"
)

func main() {
	//getting list of all the connection.
	listConnection := database.SystemConnection()

	// get mongodb connection
	mongoConn := listConnection["mongodb"].(*mgo.Session)

	// init driver model
	driverData := &driverModel.DriverData{}
	driverData.GetConn(mongoConn)

	// inserting dummy driver
	insertDummyDriver(driverData)
}

// insert database 50.000 rows
// passed driver struct to save the data to database.
func insertDummyDriver(driverData *driverModel.DriverData) {

	dummyDrivers := GenereateDriver(50000)
	for _, driver := range dummyDrivers {
		driverData.Insert(driver.Name, driver.Lat, driver.Lon, driver.Status)
	}

}

type Driver struct {
	Name   string  `json:"name"`
	Lat    float64 `json:"lat"`
	Lon    float64 `json:"lon"`
	Status bool    `json:"status"`
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
