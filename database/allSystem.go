package database

// this file is used for combining all the connection from different database system

//we create different types of databse connection here
func SystemConnection() map[string]interface{} {
	listConnection := make(map[string]interface{})
	var err error

	//create mongodb connection
	mongo := MongoHost{
		Host: "127.0.0.1",
		Port: "27017",
	}
	mongoConnection := mongo.Connect()

	listConnection["redis"] = redisConnection
	listConnection["postgre"] = postgreConnection
	listConnection["mongodb"] = mongoConnection
	return listConnection
}
