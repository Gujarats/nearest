package driverMock

import (
	"github.com/Gujarats/API-Golang/model/driver"
	mgo "gopkg.in/mgo.v2"
	redis "gopkg.in/redis.v5"
)

type DriverMongoExistMock struct{}

//===================MongoDB====================//
func (d *DriverMongoExistMock) GetConn(mongoSession *mgo.Session, redisConnection *redis.Client) {
}

func (d *DriverMongoExistMock) Insert(name string, lat, lon float64, status bool) {

}

func (d *DriverMongoExistMock) Find(name string) *driver.DriverData {
	return &driver.DriverData{}
}

func (d *DriverMongoExistMock) Update(driverData driver.DriverData) {}

// return non-empty or exist mock data
func (d *DriverMongoExistMock) GetNearLocation(distance int64, lat, lon float64) []driver.DriverData {
	return []driver.DriverData{
		{Id: "_idDummy", Name: "testDriver", Status: true, Location: driver.GeoJson{Coordinates: []float64{2.2, 2.2}}},
		{Id: "_idDummy", Name: "testDriver", Status: true, Location: driver.GeoJson{Coordinates: []float64{2.2, 2.2}}},
	}
}

//===================REDIS====================//

func (d *DriverMongoExistMock) SaveDriversRedis(drivers []driver.DriverData, city, idDistrict string) {
}

// return empty drivers from redis
func (d *DriverMongoExistMock) DriversRedis(city, idDistrict string) []driver.DriverData {
	var drivers []driver.DriverData

	return drivers
}
