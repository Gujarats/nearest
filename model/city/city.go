package city

import (
	"errors"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	redis "gopkg.in/redis.v5"
)

type City struct {
	Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name     string        `json:"name"`
	District int           `json:"district"`
	Location GeoJson       `json:"location"`
}

type GeoJson struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

// Global variable for storing database connection.
var mongo *mgo.Session
var redisConn *redis.Client

func (c *City) GetConn(mongoConnection *mgo.Session, redisConnection *redis.Client) {
	mongo = mongoConnection
	redisConn = redisConnection
}

// check mongo connection if error return it.
func checkMongoConnection(mongoConnection *mgo.Session) error {
	if mongoConnection == nil {
		return errors.New("No Mongo Connection")
	}

	return nil
}

func (c *City) CreateIndex(collectionName string) error {
	var err error
	err = checkMongoConnection(mongo)
	if err != nil {
		return err
	}

	index := mgo.Index{
		Key: []string{"$2dsphere:locatation"},
	}

	// create index from given collection
	collection := mongo.DB("Driver").C(collectionName)
	err = collection.EnsureIndex(index)
	if err != nil {
		return err
	}

	return nil

}

// Inserting district to mongo database
func (c *City) InsertDistrict(city string, distric int, lat, lon float64) error {
	var err error
	err = checkMongoConnection(mongo)
	if err != nil {
		return err
	}

	collection := mongo.DB("Driver").C(city)

	c.District = distric
	c.Name = city
	c.Location = GeoJson{Type: "point", Coordinates: []float64{lon, lat}} // lon, lat order rules from mongodb

	err = collection.Insert(c)
	if err != nil {
		return err
	}

	return nil
}

func (c *City) AllDistrict(city string) ([]City, error) {
	var cities []City
	var err error

	err = checkMongoConnection(mongo)
	if err != nil {
		return cities, err
	}

	collection := mongo.DB("Driver").C(city)
	err = collection.Find(bson.M{}).All(&cities)
	if err != nil {
		return cities, err
	}

	return cities, nil
}

func (c *City) GetNearestDistrict(cityName string, lat, lon float64, distance int64) (City, error) {
	var err error
	var city City

	err = checkMongoConnection(mongo)
	if err != nil {
		return city, err
	}

	collection := mongo.DB("Driver").C(cityName)

	err = collection.Find(bson.M{
		"location": bson.M{
			"$nearSphere": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{lon, lat}, // lon,lat in order is the rule from mongodb
				},
				"$maxDistance": distance,
			},
		},
	}).One(&city)
	if err != nil {
		return city, err
	}

	return city, nil
}
