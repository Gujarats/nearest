package driver

import (
	"github.com/training_project/util/logger"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DriverData struct {
	Name     string  `json:"name"`
	Status   bool    `json:"status"`
	Location GeoJson `json:"location"`
}

// struct for storing geo location in mongodb
type GeoJson struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

var mongo *mgo.Session

func init() {
	logger.InitLogger("Model Driver", "../../logs/", "Model.txt")
}

func (d *DriverData) GetConn(mongoSession *mgo.Session) {
	mongo = mongoSession
}

func (d *DriverData) GetNearLocation(distance int64, lat, lon float64) []DriverData {
	collection := mongo.DB("Driver").C("driver")

	var driverLocation []DriverData
	err := collection.Find(bson.M{
		"location": bson.M{
			"$nearSphere": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{lon, lat},
				},
				"$maxDistance": distance,
			},
		},
		"status": true,
	}).Limit(5).All(&driverLocation)

	if err != nil {
		logger.CheckError("Mongo", err)
	}

	return driverLocation
}

func (d *DriverData) Insert(name string, lat, lon float64, status bool) {
	// init value data
	d.Name = name
	d.Status = status
	d.Location.Type = "Point"

	// set the slice to empty,
	// to make sure there are no additional or old data exist in the slice.
	d.Location.Coordinates = []float64{}
	d.Location.Coordinates = append(d.Location.Coordinates, lon)
	d.Location.Coordinates = append(d.Location.Coordinates, lat)

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

//update data if exist if not the insert it
func (d *DriverData) Update(name string, lat, lon float64, status bool) {
	d.Name = name
	d.Status = status

	// set the slice to empty,
	// to make sure there are no additional or old data exist in the slice.
	d.Location.Coordinates = []float64{}
	d.Location.Coordinates = append(d.Location.Coordinates, lon)
	d.Location.Coordinates = append(d.Location.Coordinates, lat)

	collection := mongo.DB("Driver").C("driver")

	_, err := collection.Upsert(bson.M{"name": name}, d)

	logger.CheckError("Model Driver", err)

}
