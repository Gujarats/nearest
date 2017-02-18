package driverMock

import (
	"github.com/training_project/model/driver"
	mgo "gopkg.in/mgo.v2"
)

type DriverDataMock struct {
	Name string
}

func (d *DriverDataMock) GetConn(mongoSession *mgo.Session) {
}

func (d *DriverDataMock) Insert(name, lat, lon string, status bool) {

}

func (d *DriverDataMock) Find(name string) *driver.DriverData {
	return &driver.DriverData{}
}
