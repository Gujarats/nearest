package driverInterface

import (
	"github.com/training_project/model/driver"
	mgo "gopkg.in/mgo.v2"
)

type DriverInterfacce interface {
	GetConn(mongoSession *mgo.Session)
	Insert(name string, lat string, lon string, status bool)
	Find(name string) *driver.DriverData
	Update(name, lat, lon string, status bool)
}
