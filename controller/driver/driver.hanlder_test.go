package driver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/Gujarats/API-Golang/model/driver/interface"
	"github.com/Gujarats/API-Golang/model/driver/mock"
	"github.com/Gujarats/API-Golang/model/global"
)

func TestDriverHandlerBadRequestInputViolation(t *testing.T) {
	// create body params
	body := url.Values{}
	body.Set("name", "driver1")
	body.Set("latitude", "latExample")
	body.Set("longitude", "lonExample")
	body.Set("status", "true")

	//we pass a dummy value to pass the required params
	req := httptest.NewRequest("POST", "/driver", bytes.NewBufferString(body.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(body.Encode())))

	w := httptest.NewRecorder()

	// driver mock model
	driverDataMock := driverMock.DriverDataMock{}
	var driver driverInterface.DriverInterfacce
	driver = &driverDataMock

	handler := InsertDriver(driver)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Error actual = %v, expected = %v\n", w.Code, http.StatusBadRequest)
	}

	// check the response
	resp := w.Body.Bytes()
	if resp == nil {
		t.Error("Error Body result Empty")
	}

	RespResult := global.Response{}
	err := json.Unmarshal(resp, &RespResult)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("response = %+v\n", RespResult)
}

// Finding Driver Ok from Redis
func TestFindDriverOKRedis(t *testing.T) {
	// create body params
	body := url.Values{}
	body.Set("latitude", "48.8588377")
	body.Set("longitude", "2.2775176")
	body.Set("city", "kuningan")
	body.Set("distance", "200")

	//we pass a dummy value to pass the required params
	req := httptest.NewRequest("POST", "/driver/find", bytes.NewBufferString(body.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(body.Encode())))

	w := httptest.NewRecorder()

	// driver mock model
	driverDataMock := driverMock.DriverDataMock{}
	var driver driverInterface.DriverInterfacce
	driver = &driverDataMock

	handler := FindDriver(driver)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Error actual = %v, expected = %v\n", w.Code, http.StatusOK)
	}

	// check the response
	resp := w.Body.Bytes()
	if resp == nil {
		t.Error("Error Body result Empty")
	}

	RespResult := global.Response{}
	err := json.Unmarshal(resp, &RespResult)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("response = %+v\n", RespResult)

}

// finding driver Ok from mongo empty data
// create mock that return empty data from redis
func TestFindDriverOKMongoDB(t *testing.T) {
	// create body params
	body := url.Values{}
	body.Set("latitude", "48.8588377")
	body.Set("longitude", "2.2775176")
	body.Set("city", "kuningan")
	body.Set("distance", "200")

	//we pass a dummy value to pass the required params
	req := httptest.NewRequest("POST", "/driver/find", bytes.NewBufferString(body.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(body.Encode())))

	w := httptest.NewRecorder()

	// driver mock model
	driverDataMock := driverMock.DriverMongoMock{}
	var driver driverInterface.DriverInterfacce
	driver = &driverDataMock

	handler := FindDriver(driver)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Error actual = %v, expected = %v\n", w.Code, http.StatusOK)
	}

	// check the response
	resp := w.Body.Bytes()
	if resp == nil {
		t.Error("Error Body result Empty")
	}

	RespResult := global.Response{}
	err := json.Unmarshal(resp, &RespResult)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("response = %+v\n", RespResult)

}

// finding driver Ok from mongo exist data
// create mock that return empty data from redis
func TestFindDriverOKMongoDBExist(t *testing.T) {
	// create body params
	body := url.Values{}
	body.Set("latitude", "48.8588377")
	body.Set("longitude", "2.2775176")
	body.Set("city", "kuningan")
	body.Set("distance", "200")

	//we pass a dummy value to pass the required params
	req := httptest.NewRequest("POST", "/driver/find", bytes.NewBufferString(body.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(body.Encode())))

	w := httptest.NewRecorder()

	// driver mock model
	driverDataMock := driverMock.DriverMongoExistMock{}
	var driver driverInterface.DriverInterfacce
	driver = &driverDataMock

	handler := FindDriver(driver)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Error actual = %v, expected = %v\n", w.Code, http.StatusOK)
	}

	// check the response
	resp := w.Body.Bytes()
	if resp == nil {
		t.Error("Error Body result Empty")
	}

	RespResult := global.Response{}
	err := json.Unmarshal(resp, &RespResult)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("response = %+v\n", RespResult)

}

func TestUpdateDriverOK(t *testing.T) {
	// create body params
	body := url.Values{}
	body.Set("name", "driver1")
	body.Set("id", "DummyId")
	body.Set("latitude", "48.8588377")
	body.Set("longitude", "2.2775176")
	body.Set("status", "true")

	//we pass a dummy value to pass the required params
	req := httptest.NewRequest("POST", "/driver", bytes.NewBufferString(body.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(body.Encode())))

	w := httptest.NewRecorder()

	// driver mock model
	driverDataMock := driverMock.DriverDataMock{}
	var driver driverInterface.DriverInterfacce
	driver = &driverDataMock

	handler := UpdateDriver(driver)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Error actual = %v, expected = %v\n", w.Code, http.StatusOK)
	}

	// check the response
	resp := w.Body.Bytes()
	if resp == nil {
		t.Error("Error Body result Empty")
	}

	RespResult := global.Response{}
	err := json.Unmarshal(resp, &RespResult)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("response = %+v\n", RespResult)
}
