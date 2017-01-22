package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/training_project/database"
	"github.com/training_project/model/review"
	"github.com/training_project/model/shop"

	redis "gopkg.in/redis.v4"
	grace "gopkg.in/tokopedia/grace.v1"
)

func main() {
	//getting list of all the connection.
	listConnection := database.SystemConnection()

	//getting redis connection convert it from interface to *redisClient.
	redisConn := listConnection["redis"].(*redis.Client)

	// get postgre connection.
	postgreConn := listConnection["postgre"].(*sqlx.DB)

	//create a model object.
	activeSeller := shop.ActiveSeller{
		ShopId: 124,
	}

	//Get the connection and insert the value
	activeSeller.GetConn(redisConn)

	//insert the id
	activeSeller.InsertActiveSeller()

	//insert another id
	activeSeller.ShopId = 1
	activeSeller.InsertActiveSeller()

	//review
	review.GetConn(postgreConn)

	result := review.IsDataExist("2017-01-01", 9)
	fmt.Printf("result db = %+v\n", result)

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
