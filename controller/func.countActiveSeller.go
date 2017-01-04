package controller

import (
	"errors"
	"fmt"
	"github.com/training_project/database"
	"github.com/training_project/model"
	"log"
	"time"
)

// get active seller setbit from redis
func GetActiveSellerByte(keyActiveSeller string) []byte {
	log.SetFlags(-1)
	log.SetPrefix("redis query = ")
	result, err := database.Redisdb.Get(keyActiveSeller).Bytes()
	if err != nil {
		log.Print(err)
	}

	// print result
	log.Printf("result =%+v\n", result)

	return result
}

func insertActiveSeller(data []byte) {
	tx := database.Postgre.MustBegin()

	//get date now
	//format time now to yyy-mm-dd
	t := time.Now().Local()
	formatTime := t.Format("2006-01-02")

	tx.MustExec("INSERT INTO hello (date, active_seller) VALUES ($1, $2, $3)", formatTime, data)
	tx.Commit()
}

func countActiveSeller(date string) (int64, error) {
	//assume date is well formated
	activeSeller := []model.ActiveSeller{}
	database.Postgre.Select(&activeSeller, "SELECT * FROM hello WHERE date=$1", date)
	if len(activeSeller) == 0 {
		return 0, errors.New("DB: Active Seller Not found")
	}

	// print result
	fmt.Printf("result query =%+v\n", activeSeller)

	return 1, nil
}
