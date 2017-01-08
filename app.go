package main

import (
	"github.com/training_project/database"
	"github.com/training_project/model/shop"
	redis "gopkg.in/redis.v4"
)

func main() {
	//getting list of all the connection
	listConnection := database.SystemConnection()

	//getting redis connection convert it from interface to *redisClient
	redisConn := listConnection["redis"].(*redis.Client)

	activeSeller := model.ActiveSeller{
		ShopId: 124,
	}

	//Get the connection and insert the value
	activeSeller.GetConn(redisConn)
	activeSeller.InsertActiveSeller()

	//insert another id
	activeSeller.ShopId = 1
	activeSeller.InsertActiveSeller()
}
