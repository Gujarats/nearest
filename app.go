package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/training_project/database"
	"github.com/training_project/model/shop"
	redis "gopkg.in/redis.v4"
	grace "gopkg.in/tokopedia/grace.v1"
)

func main() {
	//getting list of all the connection
	listConnection := database.SystemConnection()

	//getting redis connection convert it from interface to *redisClient
	redisConn := listConnection["redis"].(*redis.Client)

	//create a model object
	activeSeller := model.ActiveSeller{
		ShopId: 124,
	}

	//Get the connection and insert the value
	activeSeller.GetConn(redisConn)

	//insert the id
	activeSeller.InsertActiveSeller()

	//insert another id
	activeSeller.ShopId = 1
	activeSeller.InsertActiveSeller()

	http.HandleFunc("/test", GetAllServices)
	port := "8080"
	log.Fatal(grace.Serve(fmt.Sprintf(":%s", port), nil))
}

type responseBrow struct {
	Status string `json:"status"`
}

func GetAllServices(w http.ResponseWriter, r *http.Request) {
	allServices := responseBrow{
		Status: "hello",
	}

	json, err := json.Marshal(allServices)

	if err != nil {
		log.Panic(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(json))
}
