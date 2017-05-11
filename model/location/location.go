package location

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

var mongo *mgo.Session

func Connect(mongoConn *mgo.Session) {
	mongo = mongoConn
}

// bulk insert into given collectionName
func InsertLocations(collectionName string, datas []interface{}) {
	collection := mongo.DB("Driver").C(collectionName)
	bulk := collection.Bulk()

	bulk.Insert(datas...)

	_, err := bulk.Run()
	if err != nil {
		log.Fatal(err)
	}

}
