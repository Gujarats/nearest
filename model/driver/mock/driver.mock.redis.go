package driverMock

import (
	"github.com/Gujarats/API-Golang/model/driver"
	mgo "gopkg.in/mgo.v2"
	redis "gopkg.in/redis.v5"
)

type DriverOnlyCityMock struct {
	Name string
}

//===================MongoDB====================//
func (d *DriverOnlyCityMock) GetConn(mongoSession *mgo.Session, redisConnection *redis.Client) {
}

func (d *DriverOnlyCityMock) Insert(collectionName, name string, lat, lon float64, status bool) {

}

func (d *DriverOnlyCityMock) Find(name string) *driver.DriverData {
	return &driver.DriverData{}
}

func (d *DriverOnlyCityMock) Update(driverData driver.DriverData) {}

func (d *DriverOnlyCityMock) GetNearLocation(distance int64, lat, lon float64) []driver.DriverData {
	return []driver.DriverData{}
}

// returning available driver
func (d *DriverOnlyCityMock) GetAvailableDriver() []driver.DriverData {
	var drivers []driver.DriverData

	drivers = []driver.DriverData{
		{Name: "Test"},
		{Name: "Test"},
	}
	return drivers
}

//===================REDIS====================//

func (d *DriverOnlyCityMock) SaveDriversRedis(drivers []driver.DriverData, city, idDistrict string) {
}

func (d *DriverOnlyCityMock) DriversRedis(city, idDistrict string) []driver.DriverData {
	var drivers []driver.DriverData

	return drivers
}
