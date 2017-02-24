package main

// NOTE : This file is used for integration test.
// All the test will passed assuming if all the connection to the database is ok.

import (
	"fmt"
	"testing"

	mgo "gopkg.in/mgo.v2"

	"github.com/training_project/database"
	"github.com/training_project/model/driver"
)

// struct for input test case.
type testObject struct {
	Lat          float64
	Lon          float64
	Distance     int // in meters
	ExpectedRows int
}

// list connection database
var mongoConn *mgo.Session

// list data struct
var driverData driver.DriverData

// we're going to init the connection to the database.
func init() {
	listConnection := database.SystemConnection()
	mongoConn = listConnection["mongodb"].(*mgo.Session)

	// pass the connection to struct
	driverData = driver.DriverData{}
	driverData.GetConn(mongoConn)

}

// Test case :  assuming if the database has driver's location
func TestGetNearDriver(t *testing.T) {
	testObjects := []testObject{
		{Lat: 48.8536111, Lon: 2.2993946, Distance: 300, ExpectedRows: 4},
		{Lat: 48.8536111, Lon: 2.2993946, Distance: 300, ExpectedRows: 7},
	}

	for _, test := range testObjects {
		totalDriver := driverData.GetNearLocation(test.Distance, test.Lat, test.Lon)
		fmt.Printf("Total driver found = %v\n", len(totalDriver))
		if len(totalDriver) < test.ExpectedRows {
			t.Errorf("Error : the total of result is not more than = %v\n", test.ExpectedRows)
		}
	}

}
