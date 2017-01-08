package database

// this file is used for combining all the connection from different database system

func SystemConnection() map[string]interface{} {
	listConnection := make(map[string]interface{})

	// create redis connection
	redisConn := RedisHost{
		Address:  "localhost:6379",
		Password: "",
		DB:       0,
	}

	// create postgre connection
	//postgreConn := PostgreHost{
	//	Driver:   "postgre",
	//	Database: "postgre",
	//	Username: "postgre",
	//	Ssl:      "disable",
	//	Password: "root",
	//}

	redisConnection, err := redisConn.Connect()
	if err != nil {
		panic(err)
	}

	listConnection["redis"] = redisConnection
	return listConnection
}
