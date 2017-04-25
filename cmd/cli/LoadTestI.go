package main

import mgo "gopkg.in/mgo.v2"

type LoadTestI interface {
	GetAllLoadTest(mongo *mgo.Session, collectionName string) []LoadTest
}
