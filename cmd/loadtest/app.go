package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"sync"

	mgo "gopkg.in/mgo.v2"

	"github.com/Gujarats/API-Golang/database"
	"github.com/Gujarats/API-Golang/model/global"

	location "gopkg.in/gujarats/GenerateLocation.v1"
)

type LoadTest struct {
	Latency    float64
	StatusCode int
	Message    string
	DriverName string
	DriverId   string
}

func main() {
	// get mongoDB connection
	DB := database.GetMongo()

	// base location
	lat := -6.8647721
	lon := 107.553501

	// geneerate location with distance 1 km in every point and limit lenght 50 km.
	// so it will be (50/5)^2 = 2500 district
	loc := location.New(lat, lon)
	locations := loc.GenerateLocation(0.1, 50)
	fmt.Println("length ", len(locations))
	//for _, location := range locations {
	//	fmt.Printf("location  :: %+v\n", location)
	//}

	// create channel to receive the locations
	//genLocation := genChanLocation(locations)

	// send the request
	createRequestConcurrently(DB, locations)
	//createRequestSequence(DB, locations)
}

func createRequestSequence(DB *mgo.Session, locations []location.Location) {

	for _, location := range locations {
		lat := strconv.FormatFloat(location.Lat, 'f', -1, 64)
		lon := strconv.FormatFloat(location.Lon, 'f', -1, 64)
		loadTest := findRequest("Bandung", lat, lon, "500")
		err := InsertLoadTest(DB, "loadTest1", loadTest)
		if err != nil {
			fmt.Println("Insert Data Error")
		}

	}

}

func genChanLocation(locations []location.Location) <-chan location.Location {
	result := make(chan location.Location)

	go func() {
		for _, loc := range locations {
			result <- loc
		}

		close(result)

	}()

	return result
}

func createRequestConcurrently(DB *mgo.Session, locations []location.Location) {
	var wg sync.WaitGroup

	for _, location := range locations {
		wg.Add(1)
		go func(DB *mgo.Session) {
			defer wg.Done()
			lat := strconv.FormatFloat(location.Lat, 'f', -1, 64)
			lon := strconv.FormatFloat(location.Lon, 'f', -1, 64)
			loadTest := findRequest("Bandung", lat, lon, "500")
			err := InsertLoadTest(DB, "loadTest1", loadTest)
			if err != nil {
				fmt.Println("Insert Data Error")
			}
		}(DB)

	}
	wg.Wait()
}

func findRequest(city, lat, lon, distance string) LoadTest {
	var loadTest LoadTest

	body := url.Values{}
	body.Set("latitude", lat)
	body.Set("longitude", lon)
	body.Set("city", city)
	body.Set("distance", distance)

	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://localhost:8080/driver/find", bytes.NewBufferString(body.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(body.Encode())))
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	loadTest.StatusCode = resp.StatusCode

	var result []byte
	result, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	globalResponse := global.Response{}

	err = json.Unmarshal(result, &globalResponse)
	if err != nil {
		panic(err)
	}

	// first we convert interface to a map so we can get the data from gloalResponse
	if globalResponse.Data != nil {
		data := globalResponse.Data.(map[string]interface{})
		//fmt.Printf("result = %+v\n", data)

		// get latency to struct
		loadTest.Latency = globalResponse.Latency
		loadTest.Message = globalResponse.Message

		driverName := data["name"]
		if driverName != nil {
			loadTest.DriverName = driverName.(string)
		}

		id := data["id"]
		if driverName != nil {
			loadTest.DriverId = id.(string)
		}

	}

	return loadTest

}

// Save the result to database
func InsertLoadTest(mongo *mgo.Session, collectionName string, loadTest LoadTest) error {
	collection := mongo.DB("LoadTest").C(collectionName)
	err := collection.Insert(loadTest)
	if err != nil {
		return err
	}

	return nil
}
