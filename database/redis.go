package database

import (
	"fmt"
	"gopkg.in/redis.v4"
	"time"
)

var Redisdb *redis.Client

func InsertActiveSeller(userId int64) {
	t := time.Now().Local()
	//format time now to yyy-mm-dd
	formatTime := t.Format("2006-01-02")

	keyActiveSeller := "active_seller:" + formatTime

	Redisdb.SetBit(keyActiveSeller, userId, 1)

}

func GetActiveSeller(keyActiveSeller string) {
	uniqueSeller := Redisdb.Get(keyActiveSeller)
	fmt.Print(uniqueSeller, "\n")
}

func CountUniqueUser() {

}
