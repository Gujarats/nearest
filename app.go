package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/Gujarats/nearest/database"
	"github.com/gorilla/mux"

	driverModel "github.com/Gujarats/nearest/model/driver"
	driverInterface "github.com/Gujarats/nearest/model/driver/interface"

	cityModel "github.com/Gujarats/nearest/model/city"
	cityInterface "github.com/Gujarats/nearest/model/city/interface"

	driverController "github.com/Gujarats/nearest/controller/driver"

	mgo "gopkg.in/mgo.v2"
	redis "gopkg.in/redis.v5"
)

func main() {
	//getting list of all the connection.
	listConnection := database.SystemConnection()

	//getting redis connection convert it from interface to *redisClient.
	redisConn := listConnection["redis"].(*redis.Client)

	// get postgre connection.
	mongoConn := listConnection["mongodb"].(*mgo.Session)

	// set driver interface
	var driver driverInterface.DriverInterfacce
	driverData := &driverModel.DriverData{}
	driver = driverData

	// set city interface
	var city cityInterface.CityInterfacce
	cityData := &cityModel.City{}
	city = cityData

	// pass database connections to model
	driver.GetConn(mongoConn, redisConn)
	city.GetConn(mongoConn, redisConn)
	m := &sync.Mutex{}

	// router
	router := mux.NewRouter()
	router.HandleFunc("/driver/v2/find", driverController.FindDriver(m, driver, city))
	log.Fatal(http.ListenAndServe(":8080", router))

	// driver router
	//http.Handle("/driver/v1/find", driverController.FindDriver(m, driver, city))
	http.Handle("/driver/v1/update", driverController.UpdateDriver(m, driver, city))

	port := ":8081"
	fmt.Println("App Started on port = ", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Panic("App Started Failed = ", err.Error())
	}

}
