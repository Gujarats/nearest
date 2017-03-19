package cityMock

import (
	"errors"

	mgo "gopkg.in/mgo.v2"
	redis "gopkg.in/redis.v5"
)

type CityMock struct {
}

func (c *CityMock) GetConn(mongoConnection *mgo.Session, redisConnection *redis.Client) {
}

// check mongo connection if error return it.
func checkMongoConnection(mongoConnection *mgo.Session) error {
	if mongoConnection == nil {
		return errors.New("No Mongo Connection")
	}

	return nil
}

func (c *CityMock) CreateIndex(collectionName string) error {
	return nil

}

// Inserting district to mongo database
func (c *CityMock) InsertDistrict(city string, distric int, lat, lon float64) error {
	return nil
}

func (c *CityMock) AllDistrict(city string) ([]City, error) {
	var cities []City
	var err error

	return cities, nil
}

func (c *CityMock) GetNearestDistrict(cityName string, lat, lon float64, distance int64) (City, error) {
	var err error
	var city City

	return city, nil
}
