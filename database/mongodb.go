package database

import mgo "gopkg.in/mgo.v2"

// create session variable for receiving connection mongodb
var sessionMongo *mgo.Session

type MongoHost struct {
	Host string
	Port string
}

func (m *MongoHost) Connect() *mgo.Session {
	if m.Host == "" || m.Port == "" {
		logger.Panic("Mongo Host and Port must be initialized")
	}

	session, err := mgo.Dial(m.Host + ":" + m.Port)
	if err != nil {
		logger.Panic(err)
	}

	return session
}
