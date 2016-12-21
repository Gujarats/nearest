package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/training_project/controller/handler"
	"github.com/training_project/database"
	"log"
	"net/http"
)

func main() {
	database.InitMysqlDb()
	log.Printf("App starting ...")
	router := httprouter.New()

	router.GET("/v1/talks", handler.ReadTalks)
	router.POST("/v1/talks", handler.WriteTalks)

	log.Printf("App listen on 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
