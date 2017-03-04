package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Gujarats/API-Golang/config"
	"github.com/Gujarats/API-Golang/controller/driver"
	"github.com/Gujarats/API-Golang/controller/review"
	"github.com/Gujarats/API-Golang/database"
	"github.com/jmoiron/sqlx"

	driverModel "github.com/Gujarats/API-Golang/model/driver"
	reviewModel "github.com/Gujarats/API-Golang/model/review"

	"github.com/Gujarats/API-Golang/util/logger"

	mgo "gopkg.in/mgo.v2"
)

var cfg config.Config

func init() {
	logger.InitLogger("App :: ", "./logs/", "App.txt")
}

func main() {
	//getting list of all the connection.
	listConnection := database.SystemConnection()

	//getting redis connection convert it from interface to *redisClient.
	//redisConn := listConnection["redis"].(*redis.Client)

	// get postgre connection.
	postgreConn := listConnection["postgre"].(*sqlx.DB)
	mongoConn := listConnection["mongodb"].(*mgo.Session)

	//pass connection to model to model
	reviewData := &reviewModel.ReviewData{}
	reviewData.GetConn(postgreConn)

	driverData := &driverModel.DriverData{}
	driverData.GetConn(mongoConn)

	// inserting dummy driver
	//	insertDummyDriver(driverData)

	// review router
	http.HandleFunc("/", review.CheckDataExist)

	// driver router
	http.HandleFunc("/driver", driver.InsertDriver)
	http.HandleFunc("/driver/find", driver.FindDriver)

	port := ":8080"
	fmt.Println("App Started on port = ", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Panic("App Started Failed = ", err.Error())
	}

}
