package cityMock

import (
	"errors"

	mgo "gopkg.in/mgo.v2"
	redis "gopkg.in/redis.v5"

	"github.com/Gujarats/API-Golang/model/city"
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

func (c *CityMock) AllDistrict(cityName string) ([]city.City, error) {
	var cities []city.City

	cities = []city.City{
		{Name: "Bandung", District: 1},
		{Name: "Bandung", District: 3},
	}

	return cities, nil
}

// return nil error and mock city
func (c *CityMock) GetNearestDistrict(cityName string, lat, lon float64, distance int64) (city.City, error) {
	var cityData city.City

	cityData = city.City{Name: "Bandung"}

	return cityData, nil
}
