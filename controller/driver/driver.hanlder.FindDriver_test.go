package driver

import (
	"errors"
	"net/http"
	"testing"

	"github.com/Gujarats/API-Golang/model/city"
	"github.com/Gujarats/API-Golang/model/driver"

	"github.com/Gujarats/API-Golang/model/city/mock"
	"github.com/Gujarats/API-Golang/model/driver/mock"
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
				Err: nil,
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
				Err: nil,
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
				Err: nil,
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
				Err: nil,
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
		actualStatus, actualMessage, err := createFindDriverRequest(&testObject.CityMock, &testObject.DriverMock)
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
