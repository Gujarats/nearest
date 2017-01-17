package shop

import (
	"log"
	"time"

	redis "gopkg.in/redis.v4"
)

type Connection interface {
	GetActiveSeller(*redis.Client, string) string
}

var redisConn *redis.Client

// get connection to redis and assign it to global variable unexported.
func (self *ActiveSeller) GetConn(redisConnection *redis.Client) {
	redisConn = redisConnection
}

// get active seller setbit from redis
func (self *ActiveSeller) GetActiveSeller(keyActiveSeller string) string {
	log.SetFlags(-1)
	log.SetPrefix("redis query = ")
	keyActiveSeller = "active_seller_daily:" + keyActiveSeller
	log.Println(keyActiveSeller)

	result, _ := redisConn.Get(keyActiveSeller).Result()

	return result
}

//checking the input id in redis exist or not
func (self *ActiveSeller) IsExist(shopId int64, date time.Time) bool {
	//format time now to yyy-mm-dd
	formatTime := date.Format("2006-01-02")

	keyActiveSeller := "active_seller_daily:" + formatTime

	result, err := redisConn.GetBit(keyActiveSeller, shopId).Result()

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
func (self *ActiveSeller) InsertActiveSeller() {
	//format time now to yyy-mm-dd
	t := time.Now().Local()
	formatTime := t.Format("2006-01-02")

	keyActiveSeller := "active_seller_daily:" + formatTime

	_, err := redisConn.SetBit(keyActiveSeller, self.ShopId, 1).Result()
	if err != nil {
		log.Println("Error Insert = ", err.Error())
	}

	//use expire time 3600 seconds
	seconds := 3600
	redisConn.Expire(keyActiveSeller, time.Duration(seconds)*time.Second)

}
