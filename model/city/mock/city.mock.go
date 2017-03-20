package cityMock

import (
	mgo "gopkg.in/mgo.v2"
	redis "gopkg.in/redis.v5"

	"github.com/Gujarats/API-Golang/model/city"
)

type CityMock struct {
	Err    error
	Cities []city.City
	City   city.City
}

func (c *CityMock) GetConn(mongoConnection *mgo.Session, redisConnection *redis.Client) {}

func (c *CityMock) CreateIndex(collectionName string) error {
	return c.Err
}

// Inserting district to mongo database
func (c *CityMock) InsertDistrict(city string, distric int, lat, lon float64) error {
	return c.Err
}

func (c *CityMock) AllDistrict(cityName string) ([]city.City, error) {
	return c.Cities, c.Err
}

// return nil error and mock city
func (c *CityMock) GetNearestDistrict(cityName string, lat, lon float64, distance int64) (city.City, error) {
	return c.City, c.Err
}
