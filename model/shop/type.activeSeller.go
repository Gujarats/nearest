package model

// this file use defining Active seller and its core behaviour

import (
	"log"
	"time"

	"gopkg.in/redis.v4"
)

type ActiveSeller struct {
	Date              string `db:"date"`
	shopId            string `db:"shop_id"`
	TotalActiveSeller string `db:"active_seller"`
	RedisDb           *redis.Client
}

// get active seller setbit from redis
func (self ActiveSeller) GetActiveSeller(keyActiveSeller string) string {
	log.SetFlags(-1)
	log.SetPrefix("redis query = ")
	keyActiveSeller = "active_seller_daily:" + keyActiveSeller
	log.Println(keyActiveSeller)

	result, _ := self.RedisDb.Get(keyActiveSeller).Result()

	return result
}

//checking the input id in redis exist or not
func (self ActiveSeller) IsExist(shopId int64, date time.Time) bool {
	//format time now to yyy-mm-dd
	formatTime := date.Format("2006-01-02")

	keyActiveSeller := "active_seller_daily:" + formatTime

	result, err := self.RedisDb.GetBit(keyActiveSeller, shopId).Result()

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

//insert active seller
func (self ActiveSeller) InsertActiveSeller(shopdId int64) {
	//format time now to yyy-mm-dd
	t := time.Now().Local()
	formatTime := t.Format("2006-01-02")

	keyActiveSeller := "active_seller_daily:" + formatTime

	_, err := self.RedisDb.SetBit(keyActiveSeller, userId, 1).Result()
	if err != nil {
		log.Println("Error Insert = ", err.Error())
	}

	//use expire time 3600 seconds
	seconds := 3600
	self.RedisDb.Expire(keyActiveSeller, time.Duration(seconds)*time.Second)

}
