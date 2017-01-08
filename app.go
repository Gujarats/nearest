package main

import (
	"github.com/training_project/database"
	"github.com/training_project/model/shop"
	redis "gopkg.in/redis.v4"
)

func main() {
	//getting list of all the connection
	listConnection := database.SystemConnection()

	//getting redis connection
	redisConn := listConnection["redis"].(*redis.Client)

	activeSeller := model.ActiveSeller{
		Date:   "11-01-2017",
		ShopId: 124,
	}

	activeSeller.GetConn(redisConn)
	activeSeller.InsertActiveSeller()
}
