package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/training_project/config"
	"github.com/training_project/controller/driver"
	"github.com/training_project/controller/review"
	"github.com/training_project/database"

	driverModel "github.com/training_project/model/driver"
	reviewModel "github.com/training_project/model/review"

	"github.com/training_project/util/logger"

	mgo "gopkg.in/mgo.v2"
	logging "gopkg.in/tokopedia/logging.v1"

	dummy "github.com/dummy_data/driver"
)

var cfg config.Config

func init() {
	// get config from database.ini
	// assigne to global variable cfg
	ok := logging.ReadModuleConfig(&cfg, "/etc/test", "test") || logging.ReadModuleConfig(&cfg, "config", "test")
	if !ok {
		log.Fatalln("failed to read config")
	}

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
	insertDummyDriver(driverData)

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

// insert database 50.000 rows
// passed driver struct to save the data to database.
func insertDummyDriver(driverData *driverModel.DriverData) {

	dummyDrivers := dummy.GenereateDriver(50000)
	for _, driver := range dummyDrivers {
		driverData.Insert(driver.Name, driver.Lat, driver.Lon, driver.Status)
	}

}
