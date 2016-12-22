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
	database.NewClient()
	//insert Seller 1
	database.InsertActiveSellerDaily(1)
	database.InsertActiveSellerDaily(4)
	database.InsertActiveSellerDaily(5)
	database.InsertActiveSellerDaily(8)
	database.InsertActiveSellerDaily(11)
	database.InsertActiveSellerDaily(111)
	database.InsertActiveSellerDaily(211)

	//get Active seller
	//database.GetActiveSeller("active_seller:2016-12-21")
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
