package driverInterface

import (
	"github.com/training_project/model/driver"
	mgo "gopkg.in/mgo.v2"
)

type DriverInterfacce interface {
	GetConn(mongoSession *mgo.Session)
	Insert(name string, lat float64, lon float64, status bool)
	Find(name string) *driver.DriverData
	Update(name string, lat, lon float64, status bool)
	GetNearLocation(distance int64, lat, lon float64) []driver.DriverData
}
