package driverMock

import (
	"github.com/Gujarats/API-Golang/model/driver"
	mgo "gopkg.in/mgo.v2"
	redis "gopkg.in/redis.v5"
)

type DriverDataMock struct {
	Drivers []driver.DriverData
	Driver  driver.DriverData
}

//===================MongoDB====================//
func (d *DriverDataMock) GetConn(mongoSession *mgo.Session, redisConnection *redis.Client) {
}

func (d *DriverDataMock) Insert(collecctionName, name string, lat, lon float64, status bool) {

}

func (d *DriverDataMock) Find(name string) *driver.DriverData {
	return &d.Driver
}

func (d *DriverDataMock) Update(city, idDistrict string, driver driver.DriverData) {}

func (d *DriverDataMock) GetNearLocation(distance int64, lat, lon float64) []driver.DriverData {
	return d.Drivers
}

// returning available driver
func (d *DriverDataMock) GetAvailableDriver(city, IdDistrict string) []driver.DriverData {
	return d.Drivers
}

//===================REDIS====================//

func (d *DriverDataMock) SaveDriversRedis(drivers []driver.DriverData, city, idDistrict string) {
}

func (d *DriverDataMock) DriversRedis(city, idDistrict string) []driver.DriverData {
	return d.Drivers
}
