package database

// this file is used for combining all the connection from different database system

//we create different types of databse connection here
func SystemConnection() map[string]interface{} {
	listConnection := make(map[string]interface{})

	//create mongodb connection
	mongo := MongoHost{
		Host: "172.17.0.1",
		Port: "27017",
	}
	mongoConnection := mongo.Connect()

	listConnection["mongodb"] = mongoConnection
	return listConnection
}
