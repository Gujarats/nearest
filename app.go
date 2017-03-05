package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Gujarats/API-Golang/controller/driver"
	"github.com/Gujarats/API-Golang/database"

	driverModel "github.com/Gujarats/API-Golang/model/driver"
	driverInterface "github.com/Gujarats/API-Golang/model/driver/interface"

	mgo "gopkg.in/mgo.v2"
)

func main() {
	//getting list of all the connection.
	listConnection := database.SystemConnection()

	// get postgre connection.
	mongoConn := listConnection["mongodb"].(*mgo.Session)

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
