package database

import (
	"fmt"
	"gopkg.in/redis.v4"
	"log"
	"time"
)

var Redisdb *redis.Client

func NewClient() {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	Redisdb = client
}

//checking the input id in redis exist or not
func IsIdExist(someId int64) bool {
	//format time now to yyy-mm-dd
	t := time.Now().Local()
	formatTime := t.Format("2006-01-02")

	keyActiveSeller := "active_seller_daily:" + formatTime

	result, err := Redisdb.GetBit(keyActiveSeller, someId).Result()

	if err != nil {
		log.Println("Error IsIdExist = ", err.Error())
		return false
	}

	if result == 1 {
		return true
	} else {
		return false
	}
}

func InsertActiveSellerDaily(userId int64) {
	//format time now to yyy-mm-dd
	t := time.Now().Local()
	formatTime := t.Format("2006-01-02")

	keyActiveSeller := "active_seller_daily:" + formatTime

	_, err := Redisdb.SetBit(keyActiveSeller, userId, 1).Result()
	if err != nil {
		log.Println("Error Insert = ", err.Error())
	}

}

func InsertActiveSellerWeekly(userId int64) {
	t := time.Now().Local()
	//format time now to yyy-mm-dd
	formatTime := t.Format("2006-01-02")

	keyActiveSeller := "active_seller_weekly:" + formatTime

	Redisdb.SetBit(keyActiveSeller, userId, 1)

}

func InsertActiveSellerMonthly(userId int64) {
	t := time.Now().Local()
	//format time now to yyy-mm-dd
	formatTime := t.Format("2006-01-02")

	keyActiveSeller := "active_seller_monthly:" + formatTime

	Redisdb.SetBit(keyActiveSeller, userId, 1)

}

func GetActiveSeller(keyActiveSeller string) {
	uniqueSeller := Redisdb.Get(keyActiveSeller)
	fmt.Print(uniqueSeller, "\n")
}
