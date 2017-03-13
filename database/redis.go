package database

import "gopkg.in/redis.v5"

type RedisHost struct {
	Address  string
	Password string
	DB       int
}

type RedisSytem interface {
	Connect() (*redis.Client, error)
}

// connect to Redis and return the connection if succesful and error if otherwise
func (self *RedisHost) Connect() (*redis.Client, error) {

	client := redis.NewClient(&redis.Options{
		Addr:     self.Address,
		Password: self.Password,
		DB:       self.DB,
	})

	//check redis connection
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
