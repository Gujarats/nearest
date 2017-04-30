package driver

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"net/url"
	"strconv"
	"sync"

	"github.com/Gujarats/API-Golang/model/global"

	"github.com/Gujarats/API-Golang/model/city/mock"
	"github.com/Gujarats/API-Golang/model/driver/mock"

	"github.com/Gujarats/API-Golang/model/city/interface"
	"github.com/Gujarats/API-Golang/model/driver/interface"
)

// create test request that pass all parameters requirement.
func createFindDriverRequest(cityMock *cityMock.CityMock, driverMock *driverMock.DriverDataMock) (int, string, error) {
	// create body params
	body := url.Values{}
	body.Set("latitude", "48.8588377")
	body.Set("longitude", "2.2775176")
	body.Set("city", "Paris")
	body.Set("distance", "200")

	//we pass a dummy value to pass the required params
	req := httptest.NewRequest("POST", "/driver/find", bytes.NewBufferString(body.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(body.Encode())))

	w := httptest.NewRecorder()

	// driver mock model
	var driver driverInterface.DriverInterfacce
	driver = driverMock

	// city mock model
	var city cityInterface.CityInterfacce
	city = cityMock
	m := &sync.Mutex{}

	// craete request
	handler := FindDriver(m, driver, city)
	handler.ServeHTTP(w, req)

	// check the response
	resp := w.Body.Bytes()
	if resp == nil {
		return -1, "", errors.New("Response body is empty")
	}

	// response result
	RespResult := global.Response{}
	err := json.Unmarshal(resp, &RespResult)
	if err != nil {
		return -1, "", err
	}

	return w.Code, RespResult.Message, nil

}

// create update request with pass alll the requirement parameters
func createUpdateDriverRequest(cityMock *cityMock.CityMock, driverMock *driverMock.DriverDataMock) (int, string, error) {
	// create body params
	body := url.Values{}
	body.Set("name", "driver1")
	body.Set("id", "58ccebc7ac702fc793f9e384")
	body.Set("latitude", "48.8588377")
	body.Set("longitude", "2.2775176")
	body.Set("status", "true")
	body.Set("city", "Paris")

	//we pass a dummy value to pass the required params
	req := httptest.NewRequest("POST", "/driver/update", bytes.NewBufferString(body.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(body.Encode())))

	w := httptest.NewRecorder()

	// driver mock model
	var driver driverInterface.DriverInterfacce
	driver = driverMock

	//  city mock model
	var city cityInterface.CityInterfacce
	city = cityMock
	m := &sync.Mutex{}

	// call request to get status response
	handler := UpdateDriver(m, driver, city)
	handler.ServeHTTP(w, req)

	// check the response
	resp := w.Body.Bytes()
	if resp == nil {

		return -1, "", errors.New("Response body is empty")
	}

	RespResult := global.Response{}
	err := json.Unmarshal(resp, &RespResult)
	if err != nil {
		return -1, "", err
	}

	return w.Code, RespResult.Message, nil
}
