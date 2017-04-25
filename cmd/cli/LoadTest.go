package main

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type LoadTest struct {
	Latency    float64
	StatusCode int
	Message    string
	DriverName string
	DriverId   string
}

func (l LoadTest) GetAllLoadTest(mongo *mgo.Session, collectionName string) []LoadTest {
	var loadTests []LoadTest
	collection := mongo.DB("LoadTest").C(collectionName)
	err := collection.Find(bson.M{}).All(&loadTests)
	if err != nil {
		logger.Println("Query :: ", err.Error())
	}

	return loadTests
}
