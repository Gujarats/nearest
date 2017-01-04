package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/training_project/controller/handler"
	"github.com/training_project/database"
	"log"
	"net/http"
	"time"
)

func main() {
	testRedis()
	//testCountActiveSeller()
}

func testRedis() {
	database.InitRedisDb()
	//insert Seller 1
	database.InsertActiveSellerDaily(1)
	database.InsertActiveSellerDaily(4)
	database.InsertActiveSellerDaily(5)
	database.InsertActiveSellerDaily(8)
	database.InsertActiveSellerDaily(11)
	database.InsertActiveSellerDaily(111)
	database.InsertActiveSellerDaily(211)
	database.InsertActiveSellerDaily(1211)
	database.InsertActiveSellerDaily(91211)
	database.InsertActiveSellerDaily(4294967295)
	database.InsertActiveSellerDaily(4294967296)
	database.InsertActiveSellerDaily(9294967296)

	t := time.Now().Local()
	formatTime := t.Format("2006-01-02")
	database.GetActiveSellerByte(formatTime)

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

func testCountActiveSeller() {
	// initiate databases
	database.InitPostgreDb()

	//count total active seller with specific date

	//count total active seller with specific range of date
}
