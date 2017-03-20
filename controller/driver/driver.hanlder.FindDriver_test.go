package driver

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/Gujarats/API-Golang/model/city"
	"github.com/Gujarats/API-Golang/model/driver"

	"github.com/Gujarats/API-Golang/model/city/mock"
	"github.com/Gujarats/API-Golang/model/driver/mock"

	"github.com/Gujarats/API-Golang/model/city/interface"
	"github.com/Gujarats/API-Golang/model/driver/interface"

	"github.com/Gujarats/API-Golang/model/global"
)

func TestFindDriver(t *testing.T) {
	testObjects := []struct {
		CityMock   cityMock.CityMock
		DriverMock driverMock.DriverDataMock
		Status     int
		Message    string
	}{
		// first test index
		// with all field exist city Mock
		{
			CityMock: cityMock.CityMock{
				Err: errors.New("mock error"),
				City: city.City{
					Name: "Bandung",
				},
			},

			DriverMock: driverMock.DriverDataMock{
				Drivers: []driver.DriverData{
					{Name: "test driver"},
					{Name: "test driver"},
				},
				Driver: driver.DriverData{
					Name: "test driver",
				},
			},

			Status:  http.StatusInternalServerError,
			Message: "Failed to get nearest district",
		},

		// second test index.
		// with city mock :  error nil
		{
			CityMock: cityMock.CityMock{
				Err: nil,
				City: city.City{
					Name: "Bandung",
				},
			},

			DriverMock: driverMock.DriverDataMock{
				Drivers: []driver.DriverData{
					{Name: "test driver"},
					{Name: "test driver"},
				},
				Driver: driver.DriverData{
					Name: "test driver",
				},
			},

			Status:  http.StatusOK,
			Message: "Data found",
		},

		// third test index
		// with CityMock error and cities nil
		{
			CityMock: cityMock.CityMock{
				Err: nil,
			},

			DriverMock: driverMock.DriverDataMock{
				Drivers: []driver.DriverData{
					{Name: "test driver"},
					{Name: "test driver"},
				},
				Driver: driver.DriverData{
					Name: "test driver",
				},
			},
			Status:  http.StatusInternalServerError,
			Message: "No district found",
		},

		// fourth test index
		// Drivermock complete
		{
			CityMock: cityMock.CityMock{
				Err: nil,
				City: city.City{
					Name: "Bandung",
				},
			},

			DriverMock: driverMock.DriverDataMock{
				Drivers: []driver.DriverData{
					{Name: "test driver"},
					{Name: "test driver"},
				},
			},
			Status:  http.StatusOK,
			Message: "Data found",
		},

		// fifth test index
		// Drivermock Drivers empty
		{
			CityMock: cityMock.CityMock{
				Err: nil,
				City: city.City{
					Name: "Bandung",
				},
			},

			DriverMock: driverMock.DriverDataMock{},
			Status:     http.StatusOK,
			Message:    "We couldn't find any driver",
		},
	}

	for indexTest, testObject := range testObjects {
		actualStatus, actualMessage, err := createRequest(&testObject.CityMock, &testObject.DriverMock)
		if err != nil {
			t.Error(err)
		}

		if actualStatus != testObject.Status {
			t.Errorf("Error :: index = %v, actual = %v, expected = %v", indexTest, actualStatus, testObject.Status)
		}

		if actualMessage != testObject.Message {
			t.Errorf("Error :: index = %v, actual = %v, expected = %v", indexTest, actualMessage, testObject.Message)

		}
	}
}

// create test request that pass all parameters requirement.
func createRequest(cityMock *cityMock.CityMock, driverMock *driverMock.DriverDataMock) (int, string, error) {
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
	var driver driverInterface.DriverInterfacce
	driver = driverMock

	// city mock model
	var city cityInterface.CityInterfacce
	city = cityMock

	// craete request
	handler := FindDriver(driver, city)
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
