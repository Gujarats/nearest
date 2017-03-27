package database

import (
	"log"
	"os"
)

// this file is used for combining all the connection from different database system.

// create logger to print error in the console
var logger *log.Logger

func init() {
	logger = log.New(os.Stderr,
		"Controller Driver :: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

//we create different types of databse connection here.
func SystemConnection() map[string]interface{} {
	listConnection := make(map[string]interface{})
	var err error
	// create redis connection
	redisConn := RedisHost{
		Address:  "redis:6379",
		Password: "",
		DB:       0,
	}

	redisConnection, err := redisConn.Connect()
	if err != nil {
		panic(err)
	}

	//create mongodb connection
	mongo := MongoHost{
		Host: "db",
		Port: "27017",
	}
	mongoConnection := mongo.Connect()

	listConnection["redis"] = redisConnection
	listConnection["mongodb"] = mongoConnection
	return listConnection
}
