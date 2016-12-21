package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/training_project/controller/handler"
	"github.com/training_project/database"
	"log"
	"net/http"
)

func main() {
	testRedis()
}

func testRedis() {
	database.ExampleNewClient()
	database.Testing()
	//insert Seller 1
	database.InsertActiveSeller(1)
	database.InsertActiveSeller(2)
	database.InsertActiveSeller(3)
	database.InsertActiveSeller(4)
	database.InsertActiveSeller(5)
	database.InsertActiveSeller(6)
	database.InsertActiveSeller(8)

	//get Active seller
	database.GetActiveSeller("active_seller:2016-12-21")
}

func testAPI() {
	database.InitMysqlDb()
	log.Printf("App starting ...")
	router := httprouter.New()

	router.GET("/v1/talks", handler.ReadTalks)
	router.POST("/v1/talks", handler.WriteTalks)

	log.Printf("App listen on 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
