package database

import (
	"fmt"
	"gopkg.in/redis.v4"
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

func TestGetValue(formatTime string) {
	val, err := Redisdb.Get("active_seller_daily:" + formatTime).Result()
	if err != nil {
		fmt.Printf("error = %s\n", err)
		panic(err)
	}
	fmt.Println("key", val)
}

func InsertActiveSellerDaily(userId int64) {
	t := time.Now().Local()
	//format time now to yyy-mm-dd
	formatTime := t.Format("2006-01-02")

	keyActiveSeller := "active_seller_daily:" + formatTime

	Redisdb.SetBit(keyActiveSeller, userId, 1)

	TestGetValue(formatTime)
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
