
# API-Golang [![Build Status](https://secure.travis-ci.org/Gujarats/API-Golang.png)](http://travis-ci.org/Gujarats/API-Golang)
This is simple project for getting a driver like for example if you're using uber or grab or gojek to get a driver.

Imagine if there are 2.500.000 drivers moving around in the city. let say in Bandung city.
So if we created save those drivers in one table it is going to be a nightmare if you wanted to query the location of the drivers that nearest to user location. Why? because it is going to takes a lof of time to get only a driver that near to your location in more than 2 million of rows.

# Tech Stack

* Mongodb

* Redis

* Go

No-SQL is very easy to modify whenever it comes to change or add the field of an object stored in the collections. And we can easily create another collections to make new documents to strore the new object. this flexible database is well suited for storing the location and drivers in the city. Because the collections and the data may varies in every each city. It is my personal opinion if you have yours and better on performance then I would love to hear it.

# How to run this ?

If you are don't want to install all the databases and the dependencies. Then [docker-compose](https://docs.docker.com/compose/install/) is your answer, please install it first.
and then clonse my [compose file](https://github.com/Gujarats/Go-docker) project.

After that `cd Go-docker` inside the Go-docker folder clone this project 
```shell
$ git clone https://github.com/Gujarats/API-Golang
```

Now after all the project are cloned. the only thing you have to do to run this project is run the docker-compose command in Go-docker folder :
```shell
$ docker-compose up
```
Now you can see the project running on your machine. If you choose to not use docker-compose then you can install all the databases and dependencies all manually.

## Seed Database
Now if you have done install all Tech stack then you have to seed the database in this case in mongodb, since you have empty data in it.
It is very simple just follow this steps :
```shell
$ cd API-Golang/cmd/dummy
$ go build
$ ./dummy
```
Now all the data is seeded please wait for couple of minutes :D. After all done you can test the API and see the data collection in mongodb.
### Using Docker
If you're using docker to run this project then you need to go to bash in the running docker
```shell
$ docker exec -it YOUR_ID_CONTAINER bash
```
And then run the binary file using `./dummy`.

## Data structure
So I created some mark in the city, and saved those mark in the database. I use no-SQL to achieve this. These mark maybe varies in every city, but in this simple project I mark the city for about `2500` location and it is generated from the algorithm in the `cmd/dummy`.

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
Allright this one document will be generated in `Bandung` collection. You can always insert new mark location manually if you want. but this would be generated using `dummy data` in the `cmd/dummy` folder.


After the mark location is generated the dummy drivers will be generated to every each marked location. I will insert 1000 drivers on each marked location. So if we have 2500 marked location that would genereate 2500.000 drivers document.

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

#### Indexing
Indexing is the important part here in order to gain speed for read performance. For getting the nearest marked location in the city from given latitude and longitude I used `2dsphere` to index the `location` field in order to increase reading speed. For the drivers document the indexes field are `status` and `location`.

## Prerequisite

* Install Mongodb.

* Install Redis.

### configuration 
Go to `database` folder and config the host and port for `mongodb` and `redis`. As default it is run on default configuration for localhost. You can change it on these file : 

* mongodb.go 

* redis.go

<b>This is really important</b>
After all database connections are set. The database should be seeded using dummy datas to mongodb.
go to `cmd/dummy` and build it using `go build` and then run the program using `./dummy` command.
This will generate a marked location in Bandung and dummy drivers in every each marked location and also with the indexes field.  
The base location is : 

* latitude = -6.8647721, longitude = 107.553501

## Structure Folder
* `model `: this is where I store database transaction like read update delete etc.

    * interface : define the interface of the struct.

    * mock : define the mock struct for unit test purposes.

    * file.go : all logic query the "file" named accordingly with the root folder.
    
    * And there is `global` package for define the global response on API

* `cmd `: another main program to create dummy data.

* `controller `: this is where business logic or handler to handle the incoming request.Inside this folder there will be multiple folders and each folder will responsible for the hanlder and bussiness logic in this application.

* `database `: all database connection define here.

* `util `: all the function that will be used in all others package.


## Flow code 
In app.go : call all the connection databases-> pass them to models -> call hanlder for hanlde the incoming request.
Basically the flow look like this : main->database-> model -> handler->controller

## Load testing
I'm using [Vegeta](https://github.com/tsenart/vegeta) tools for load testing. I test the update and find API, the first test happens is find API for 5 seconds and the after that update API. here the command to do the testing.
for update API
```shell
$ echo "GET http://localhost:8080/driver/update?&city=Bandung&latitude=-6.978690151910177&longitude=108.12333333333333&distance=100&id=58ccebc7ac702fc793f9e384&status=true&name=asdf" | vegeta attack -duration=60s -rate=50 | tee results4.bin | vegeta report
```
for find API
```shell
$ echo "GET http://localhost:8080/driver/find?&city=Bandung&latitude=-6.978690151910177&longitude=108.12333333333333&distance=100" | vegeta attack -duration=60s -rate=50 | tee results_find_driver.bin | vegeta report
```
From above two commands I careated 50 request per second in 60 second.
![result_load_testing_from_local](https://github.com/Gujarats/API-Golang/blob/master/result_pike_local.png)




