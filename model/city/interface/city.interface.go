package cityInterface

import (
	mgo "gopkg.in/mgo.v2"
	redis "gopkg.in/redis.v5"

	"github.com/Gujarats/API-Golang/model/city"
)

type CityInterfacce interface {
	GetConn(mongoConnection *mgo.Session, redisConnection *redis.Client)

	// check mongo connection if error return it.
	checkMongoConnection(mongoConnection *mgo.Session) error

	CreateIndex(collectionName string) error

	// Inserting district to mongo database
	InsertDistrict(city string, distric int, lat, lon float64) error

	AllDistrict(city string) ([]city.City, error)

	GetNearestDistrict(cityName string, lat, lon float64, distance int64) (city.City, error)
}
