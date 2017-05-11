package driver

import (
	"encoding/json"
	"log"
	"os"
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
		"Driver Model :: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func (d *DriverData) GetConn(mongoSession *mgo.Session, redisConnection *redis.Client) {
	mongo = mongoSession
	redisConn = redisConnection
}

//===================REDIS====================//

// the format of the key is : city_district_id-mongodb
// saves drivers data to redis
func (d *DriverData) SaveDriversRedis(drivers []DriverData, city, idDistrict string) {
	byteDrivers, _ := json.Marshal(drivers)

	// the key here is city
	key := getFormatDistrict(city, idDistrict)
	redisConn.Set(key, byteDrivers, 0)
}

// the format of the key is : city_district_id-mongodb
// return the Drivers data from redis
func (d *DriverData) DriversRedis(city, idDistrict string) []DriverData {
	key := getFormatDistrict(city, idDistrict)
	var drivers []DriverData

	driversBytes, err := redisConn.Get(key).Bytes()
	if err != nil {
		logger.Println(err)
		return drivers
	}

	// checkking the result data from redis
	if len(driversBytes) == 0 {
		logger.Println("Drivers in redis nil")
		return drivers
	}

	// unmarshal driversBytes
	err = json.Unmarshal(driversBytes, &drivers)
	if err != nil {
		logger.Println(err)
	}

	return drivers
}

// saving the last location drivers in redis.
// the purpose is so that we can make sure the drivers data is not exist in the last collection,
// if in case the drivers go to new dristrict and update his status in new district.
func (d *DriverData) SaveLastDistrict(idDriver, city, idDistrict string) {

	key := getDriverFormatkey(idDriver)

	data := getFormatDistrict(city, idDistrict)

	dateNow := time.Now().Format("02-01-2006")

	// set expire time at midnight
	expireTime := dateNow + " 00:00"

	setTime, err := time.Parse("02-01-2006 15:04", expireTime)
	if err != nil {
		logger.Println(err)
	}

	redisConn.Set(key, data, 0)
	redisConn.ExpireAt(key, setTime)

}

// to get driver last location we used the date and their unique id from redis
// and get the
func (d *DriverData) GetLastDistrict(idDriver string) string {

	key := getDriverFormatkey(idDriver)

	result, err := redisConn.Get(key).Result()
	if err != nil {
		return ""
	}

	return result

}

//===================MongoDB====================//

// get Near driver with given distance in meters
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

func (d *DriverData) GetAvailableDriver(city, idDistrict string) []DriverData {
	collectionKey := getFormatDistrict(city, idDistrict)

	collection := mongo.DB("Driver").C(collectionKey)

	var drivers []DriverData
	err := collection.Find(bson.M{
		"status": true,
	}).Limit(100).All(&drivers)

	if err != nil {
		logger.Println(err)
	}

	return drivers
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

func (d *DriverData) InsertBulk(collectionName string, drivers []interface{}) error {
	collection := mongo.DB("Driver").C(collectionName)
	bulk := collection.Bulk()
	bulk.Insert(drivers...)

	_, err := bulk.Run()
	if err != nil {
		return err
	}

	return nil
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
func (d *DriverData) Update(city, idDistrict string, driver DriverData) error {
	collectionKey := getFormatDistrict(city, idDistrict)
	collection := mongo.DB("Driver").C(collectionKey)

	_, err := collection.Upsert(bson.M{"_id": driver.Id}, driver)
	if err != nil {
		logger.Println(err)
		return err
	}

	return nil
}

func (d *DriverData) Remove(idDriver, collectionKey string) {
	collection := mongo.DB("Driver").C(collectionKey)
	err := collection.Remove(
		bson.M{
			"_id": bson.ObjectIdHex(idDriver),
		},
	)

	if err != nil {
		logger.Println(err)
	}

}

// ============= PRIVATE FUNCTION ============= //

// return format district it is for the naming for the collections in every marked location in the city.
func getFormatDistrict(city, idDistrict string) string {
	return city + "_district_" + idDistrict
}

// return driver key format for redis key
func getDriverFormatkey(idDriver string) string {
	dateNow := time.Now().Format("02-01-2006")
	return idDriver + "_" + dateNow
}
