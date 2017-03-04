package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Gujarats/API-Golang/config"
	"github.com/Gujarats/API-Golang/controller/driver"
	"github.com/Gujarats/API-Golang/database"
	"github.com/jmoiron/sqlx"

	driverModel "github.com/Gujarats/API-Golang/model/driver"
	driverInterface "github.com/Gujarats/API-Golang/model/driver/interface"
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

	// set driver interface
	var driverInterface driverInterface.DriverInterfacce
	driverData := &driverModel.DriverData{}
	driverInterface = driverData

	driverInterface.GetConn(mongoConn)

	// driver router
	http.Handle("/driver/find", driver.FindDriver(driverInterface))
	http.Handle("/driver/update", driver.UpdateDriver(driverInterface))

	port := ":8080"
	fmt.Println("App Started on port = ", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Panic("App Started Failed = ", err.Error())
	}

}
