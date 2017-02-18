package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/training_project/config"
	"github.com/training_project/controller/review"
	"github.com/training_project/database"
	reviewModel "github.com/training_project/model/review"
	"github.com/training_project/util/logger"

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

	logger.InitLogger("App :: ", "./logs", "App.txt")
}

func main() {
	//getting list of all the connection.
	listConnection := database.SystemConnection()

	//getting redis connection convert it from interface to *redisClient.
	//redisConn := listConnection["redis"].(*redis.Client)

	// get postgre connection.
	postgreConn := listConnection["postgre"].(*sqlx.DB)

	//pass to model
	reviewData := &reviewModel.ReviewData{}
	reviewData.GetConn(postgreConn)

	http.HandleFunc("/", review.CheckDataExist)
	http.HandleFunc("/driver", review.CheckDataExist)

	port := ":8080"
	fmt.Println("App Started on port = ", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Panic("App Started Failed = ", err.Error())
	}

}
