package database

// this file is used for combining all the connection from different database system

//we create different types of databse connection here
func SystemConnection() map[string]interface{} {
	listConnection := make(map[string]interface{})
	var err error
	// create redis connection
	redisConn := RedisHost{
		Address:  "localhost:6379",
		Password: "",
		DB:       0,
	}

	redisConnection, err := redisConn.Connect()
	if err != nil {
		panic(err)
	}

	// create postgre connection
	postgreConn := PostgreHost{
		Driver:   "postgres",
		Database: "postgres",
		Username: "postgres",
		Ssl:      "disable",
		Password: "root",
	}

	postgreConnection, err := GetPostgreDb(&postgreConn)

	if err != nil {
		panic(err)
	}

	listConnection["redis"] = redisConnection
	listConnection["postgre"] = postgreConnection
	return listConnection
}
