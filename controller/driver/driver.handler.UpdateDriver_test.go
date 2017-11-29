package driver

import (
	"errors"
	"net/http"
	"testing"

	"github.com/Gujarats/nearest/model/city"
	"github.com/Gujarats/nearest/model/driver"

	"github.com/Gujarats/nearest/model/city/mock"
	"github.com/Gujarats/nearest/model/driver/mock"
)

func TestUpdateDriver(t *testing.T) {
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
	}

	for indexTest, testObject := range testObjects {
		actualStatus, actualMessage, err := createUpdateDriverRequest(&testObject.CityMock, &testObject.DriverMock)
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
