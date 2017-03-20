package driverInterface

import (
	"github.com/Gujarats/API-Golang/model/driver"
	mgo "gopkg.in/mgo.v2"
	redis "gopkg.in/redis.v5"
)

type DriverInterfacce interface {
	GetConn(mongoSession *mgo.Session, redsiConn *redis.Client)
	Insert(collectionName string, name string, lat, lon float64, status bool)
	Find(name string) *driver.DriverData

	Update(city, idDistrict string, driverData driver.DriverData)

	GetNearLocation(distance int64, lat, lon float64) []driver.DriverData
	GetAvailableDriver(city, IdDistrict string) []driver.DriverData

	SaveDriversRedis(drivers []driver.DriverData, city, idDistrict string)
	DriversRedis(city, idDistrict string) []driver.DriverData
}
