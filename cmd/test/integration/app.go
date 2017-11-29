package main

// NOTE : This file is used for integration test. Testing the API endpoint and see if the result has no error and behave like expected.

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/Gujarats/nearest/model/driver"
	"github.com/Gujarats/nearest/model/global"
)

func main() {
	testCase1()
}

// user find a driver
// after a while driver update his status with new location
func testCase1() {
	// user location
	lat := "-7.148333333333333"
	lon := "108.08891584397783"
	city := "Bandung"
	distance := "500"

	driverData := findRequest(city, lat, lon, distance)
	time.Sleep(1 * time.Second)
	fmt.Println("Success find Driver Request !!!")
	fmt.Println("===============================")
	fmt.Println("Starting update driver request")

	updateRequest("Bandung", driverData)

}

func findRequest(city, lat, lon, distance string) driver.DriverData {
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
	if resp.StatusCode != http.StatusOK {
		panic(errors.New("Response status code is not 200."))
	}

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

	driverData := driver.DriverData{}

	// first we convert interface to a map so we can get the data from gloalResponse
	data := globalResponse.Data.(map[string]interface{})
	fmt.Printf("data = %+v\n", data)
	// set all data to the specific field
	idDriver := data["id"].(string)
	driverData.Id = bson.ObjectIdHex(idDriver)

	driverData.Name = data["name"].(string)
	driverData.Status = data["status"].(bool)

	location := data["location"].(map[string]interface{})

	typeLocation := location["type"].(string)
	coordinates := location["coordinates"].([]interface{})
	coordiates1 := coordinates[0].(float64)
	coordiates2 := coordinates[1].(float64)

	driverData.Location = driver.GeoJson{Type: typeLocation, Coordinates: []float64{coordiates1, coordiates2}}

	fmt.Printf("driver Data= %+v\n", driverData)

	return driverData

}

func updateRequest(city string, driverData driver.DriverData) global.Response {

	name := driverData.Name
	lat := strconv.FormatFloat(driverData.Location.Coordinates[1], 'f', -1, 64)
	lon := strconv.FormatFloat(driverData.Location.Coordinates[0], 'f', -1, 64)
	status := strconv.FormatBool(driverData.Status)
	idDriver := driverData.Id.Hex()

	body := url.Values{}
	body.Set("name", name)
	body.Set("id", idDriver)
	body.Set("latitude", lat)
	body.Set("longitude", lon)
	body.Set("status", status)
	body.Set("city", city)

	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://localhost:8080/driver/update", bytes.NewBufferString(body.Encode()))
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
	if resp.StatusCode != http.StatusOK {
		//panic(errors.New("Response is not 200"))
	}

	var result []byte
	result, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	resposnseResult := global.Response{}

	json.Unmarshal(result, &resposnseResult)
	fmt.Printf("Update result = %+v\n", resposnseResult)

	return resposnseResult

}
