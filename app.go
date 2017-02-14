package main

import (
	"log"

	"github.com/training_project/config"
	"github.com/training_project/database"
	"github.com/training_project/model/shop"

	redis "gopkg.in/redis.v4"
	logging "gopkg.in/tokopedia/logging.v1"
)

var cfg config.Config

func init() {
	// get config from database.ini
	// assigne to global variable cfg
	ok := logging.ReadModuleConfig(&cfg, "/etc/test", "test") || logging.ReadModuleConfig(&cfg, "config", "test")
	if !ok {
		log.Fatalln("failed to read config")
	}
}

func main() {
	//getting list of all the connection.
	listConnection := database.SystemConnection()

	//getting redis connection convert it from interface to *redisClient.
	redisConn := listConnection["redis"].(*redis.Client)

	// get postgre connection.
	//postgreConn := listConnection["postgre"].(*sqlx.DB)

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
	//review.GetConn(postgreConn)

	//result := review.IsDataExist("2017-01-01", 9)
	//fmt.Printf("result db = %+v\n", result)

}
