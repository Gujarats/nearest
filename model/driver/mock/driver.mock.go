package driverMock

import (
	"github.com/Gujarats/API-Golang/model/driver"
	mgo "gopkg.in/mgo.v2"
	redis "gopkg.in/redis.v5"
)

type DriverDataMock struct {
	Name string
}

//===================MongoDB====================//
func (d *DriverDataMock) GetConn(mongoSession *mgo.Session, redisConnection *redis.Client) {
}

func (d *DriverDataMock) Insert(name string, lat, lon float64, status bool) {

}

func (d *DriverDataMock) Find(name string) *driver.DriverData {
	return &driver.DriverData{}
}

func (d *DriverDataMock) Update(driverData driver.DriverData) {}

func (d *DriverDataMock) GetNearLocation(distance int64, lat, lon float64) []driver.DriverData {
	return []driver.DriverData{}
}

//===================REDIS====================//
func (d *DriverDataMock) SaveDriversRedis(drivers []driver.DriverData, city driver.City) {
}

func (d *DriverDataMock) DriversRedis(key string) (driver.City, []driver.DriverData) {

	city := driver.City{Name: "Test"}
	drivers := []driver.DriverData{
		{Name: "TestDriversName", Status: true},
		{Name: "TestDriversName", Status: true},
	}

	return city, drivers
}
