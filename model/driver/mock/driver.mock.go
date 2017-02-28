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

func (d *DriverDataMock) Insert(name string, lat, lon float64, status bool) {

}

func (d *DriverDataMock) Find(name string) *driver.DriverData {
	return &driver.DriverData{}
}

func (d *DriverDataMock) Update(name string, lat, lon float64, status bool) {}

func (d *DriverDataMock) GetNearLocation(distance int64, lat, lon float64) []driver.DriverData {
	return []driver.DriverData{}
}
