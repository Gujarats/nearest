package driverMock

import (
	"github.com/Gujarats/API-Golang/model/driver"
	mgo "gopkg.in/mgo.v2"
	redis "gopkg.in/redis.v5"
)

type DriverMongoMock struct{}

//===================MongoDB====================//
func (d *DriverMongoMock) GetConn(mongoSession *mgo.Session, redisConnection *redis.Client) {
}

func (d *DriverMongoMock) Insert(name string, lat, lon float64, status bool) {

}

func (d *DriverMongoMock) Find(name string) *driver.DriverData {
	return &driver.DriverData{}
}

func (d *DriverMongoMock) Update(driverData driver.DriverData) {}

func (d *DriverMongoMock) GetNearLocation(distance int64, lat, lon float64) []driver.DriverData {
	return []driver.DriverData{}
}

//===================REDIS====================//
func (d *DriverMongoMock) SaveDriversRedis(drivers []driver.DriverData, city, idDistrict string) {
}

// return an empty drivers from redis
func (d *DriverMongoMock) DriversRedis(city, idDistrict string) []driver.DriverData {
	var drivers []driver.DriverData

	return drivers
}
