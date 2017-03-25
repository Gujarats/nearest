package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Gujarats/API-Golang/database"

	driverModel "github.com/Gujarats/API-Golang/model/driver"
	driverInterface "github.com/Gujarats/API-Golang/model/driver/interface"

	cityModel "github.com/Gujarats/API-Golang/model/city"
	"github.com/Gujarats/API-Golang/model/city/interface"

	driverController "github.com/Gujarats/API-Golang/controller/driver"

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

	// driver router
	http.Handle("/driver/find", driverController.FindDriver(driver, city))
	http.Handle("/driver/update", driverController.UpdateDriver(driver, city))

	port := ":8080"
	fmt.Println("App Started on port = ", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Panic("App Started Failed = ", err.Error())
	}

}
