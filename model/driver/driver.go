package driver

import (
	"encoding/json"
	"log"
	"os"
	"strings"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	redis "gopkg.in/redis.v5"
)

type DriverData struct {
	Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name     string        `json:"name"`
	Status   bool          `json:"status"`
	Location GeoJson       `json:"location"`
}

// struct for storing geo location in mongodb
type GeoJson struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type City struct {
	Name string  `json:"name"`
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
}

// Global variable for storing database connection.
var mongo *mgo.Session
var redisConn *redis.Client
var logger *log.Logger

func init() {
	logger = log.New(os.Stderr,
		"Driver",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func (d *DriverData) GetConn(mongoSession *mgo.Session, redisConnection *redis.Client) {
	mongo = mongoSession
	redisConn = redisConnection
}

//===================REDIS====================//
func (d *DriverData) SaveDriversRedis(drivers []DriverData, city City) {
	savedTime := time.Now().Local().Format("01-02-2016")
	byteDrivers, _ := json.Marshal(drivers)
	byteCity, _ := json.Marshal(city)
	data := string(byteCity) + "++" + string(byteDrivers)

	// the key here is city;date;number
	redisConn.Set(city.Name+";"+savedTime, data, 0)
}

// the format of the key is : city;date
// return the Drivers data from redis
func (d *DriverData) DriversRedis(key string) (City, []DriverData) {
	var city City
	var drivers []DriverData

	dataString, _ := redisConn.Get(key).Result()
	if dataString == "" {
		logger.Println("Drivers in redis nil")
		return city, drivers
	}

	// split data with ++
	dataSplit := strings.Split(dataString, "++")
	byteCity := []byte(dataSplit[0])
	byteDrivers := []byte(dataSplit[1])

	err := json.Unmarshal(byteCity, &city)
	if err != nil {
		logger.Println(err)
	}
	err = json.Unmarshal(byteDrivers, &drivers)
	if err != nil {
		logger.Println(err)
	}

	return city, drivers
}

//===================MongoDB====================//

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
	}).Limit(100).All(&driverLocation)

	if err != nil {
		logger.Println(err)
	}

	return driverLocation
}

func (d *DriverData) Insert(collectionName string, name string, lat, lon float64, status bool) {
	// init value data
	d.Name = name
	d.Status = status
	d.Location.Type = "Point"

	// set the slice to empty,
	// to make sure there are no additional or old data exist in the slice.
	d.Location.Coordinates = []float64{lon, lat}

	collection := mongo.DB("Driver").C(collectionName)

	err := collection.Insert(d)
	if err != nil {
		logger.Println(err)
	}

}

// create index for location and status for speed read query.
func (d *DriverData) CreateIndex(collectionName string) error {

	collection := mongo.DB("Driver").C(collectionName)
	index := mgo.Index{
		Key: []string{
			"$2dsphere:location",
			"status",
		},
	}
	err := collection.EnsureIndex(index)
	if err != nil {

		return err
	}

	return nil
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
func (d *DriverData) Update(driver DriverData) {
	collection := mongo.DB("Driver").C("driver")

	_, err := collection.Upsert(bson.M{"_id": driver.Id}, driver)

	logger.Println(err)

}
