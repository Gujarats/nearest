
# API-Golang [![Build Status](https://secure.travis-ci.org/Gujarats/API-Golang.png)](http://travis-ci.org/Gujarats/API-Golang)
This is simple project for getting a driver like for example if you're using uber or grab or gojek to get a driver.

Imagine if there are 2.500.000 drivers moving around in the city. let say in Bandung city.
So if we created save those drivers in one table it is going to be a nightmare if you wanted to query the location of the drivers that nearest to user location. Why? because it is going to takes a lof of time to get only a driver in more than 2 million of rows. <br>

####

So I created some mark in the city, and save those mark in the database. I use no-SQL and Redis to achieve this. These mark maybe varies in every city, but in this simple project I mark `1000` location in Bandung city.

The marking location algorithm is simple. Get based latitude and longitude from the edge of the city, in this case west and north. And then genereate a mark location from that base location to the east and south, those location will be separate from given distance.

I used <b>mongodb</b> here for the no-SQL database,but you can choose other no-SQL database if it is better choice. this is one document will look like in the `Bandung` collection : 

```
{
	"_id" : ObjectId("58cce7bbac702fc793bfbd77"), // this is id automatically generated from mongodb

	"name" : "Bandung",
	"district" : 0, // this district number is the naming or the mark location 
	"location" : {
		"type" : "Point",
		"coordinates" : [
			107.56489280519102,
			-6.8647721
		]
	}
}
```

Allright this one document will be generated in `Bandung` collection for more than 1000 documents for marking the location. You can always insert new mark location manually if you want. but this would be generated using `dummy data` in the `cmd/dummy` folder.

And for getting the nearest location from given latitude and longitude I used `2dsphere` and creating the index to increase reading speed.

After 1000 of mark location is generated. Now it is time to insert a dummy driver to every each marked location or district. I will insert 1000 drivers in every district. so if we have 2500 marked location that would genereate 2500.000 drivers document.

A driver document in the other hand will look like this not really different from above document : 
```
{
	"_id" : ObjectId("58ccececac702fc7930ae1ff"), // this is id automatically generated from mongodb
	"name" : "Marilyn Hicks",
	"status" : true,
	"location" : {
		"type" : "Point",
		"coordinates" : [
			108.12444444444445,
			-6.933122931146107
		]
	}
}
```
Also with indexes for the location and status important due to read speed.

####








## Structure Folder
--> model : this is where I store database transaction like read update delete etc.<br><br>
--> controller : this is where business logic or handler to handle the incoming request.<br><br>
--> database : all database connection define here.<br><br>
--> logs : this folder is used for storing the log files.<br><br>
--> util : all the function that will be used in all others package.<br><br>

## Prerequisite
    * Install Mongodb.
    * Install Redis.

### configuration 
Go to `database` folder and config the host and port for `mongodb` and `redis`. As default it is run on default configuration for localhost.
    You can change it on these file : 
    * mongodb.go 
    * redis.go

<b>This is really important</b>
After all database connections are set. Now it is time to seed a dummy data to mongodb.
go to `cmd/dummy` and build it using `go build` and then run the program using `./dummy` command.
This will generate a dummy location in Bandung :
	* latitude = -6.8647721
	* longitude = 107.553501

### Note Generating location database.
the base location 

## Code Structure
<b> Model </b>
    - interface : define the interface of the struct.
    - mock : define the mock struct for unit test purposes.
And there is `global` package for define the global response on API


<b> Util </b>
This package is used to declare the global function that used for all packages.


<b> Database</b>
This package is to define all the database connection.

<b> Controller </b>
this package is used for declare the business logic and the hanlder for handling the incoming request.

<b>Flow code </b>
In app.go : call all the connection database (database) -> pass it to model (model) -> call hanlder for hanlde the incoming request (controller)
Basically liket this : main->database-> model -> handler->controller


