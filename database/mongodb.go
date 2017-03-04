package database

import (
	"github.com/Gujarats/API-Golang/util/logger"
	mgo "gopkg.in/mgo.v2"
)

// create session variable for receiving connection mongodb
var sessionMongo *mgo.Session

func init() {
	logger.InitLogger("MONGODB :: ", "../logs", "Mongo.txt")
}

type MongoHost struct {
	Host string
	Port string
}

func (m *MongoHost) Connect() *mgo.Session {
	if m.Host == "" || m.Port == "" {
		logger.FatalLog("Mongo Host and Port must be initialized")
	}

	session, err := mgo.Dial(m.Host + ":" + m.Port)
	logger.CheckError("MongoDb", err)

	return session
}
