package driver

import (
	"github.com/training_project/util/logger"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DriverData struct {
	Name   string
	Lat    string
	Lon    string
	Status bool
}

var mongo *mgo.Session

func init() {
	logger.InitLogger("Model Driver", "../../logs", "Model.txt")
}

func (d *DriverData) GetConn(mongoSession *mgo.Session) {
	mongo = mongoSession
}

func (d *DriverData) Insert(name, lat, lon string, status bool) {
	// init value data
	d.Name = name
	d.Lat = lat
	d.Lon = lon
	d.Status = status

	collection := mongo.DB("Driver").C("driver")

	err := collection.Insert(d)

	logger.CheckError("Model Driver", err)

}

func (d *DriverData) Find(name string) *DriverData {
	collection := mongo.DB("Driver").C("driver")

	err := collection.Find(bson.M{"name": name}).One(d)
	// return empy struct if err
	if err != nil {
		return &DriverData{}
	}

	return d
}
