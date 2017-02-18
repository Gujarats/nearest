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

	"github.com/training_project/model/driver/instance"
	"github.com/training_project/model/driver/mock"
	"github.com/training_project/model/global"
)

func TestDriverHandlerOK(t *testing.T) {
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

	driverMock := driverMock.DriverDataMock{}
	driverInstance.SetInstance(driverMock)

	InsertDriver(w, req)

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

func TestFindDriverOK(t *testing.T) {
	// create body params
	body := url.Values{}
	body.Set("name", "driver1")

	//we pass a dummy value to pass the required params
	req := httptest.NewRequest("POST", "/driver/find", bytes.NewBufferString(body.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(body.Encode())))

	w := httptest.NewRecorder()

	driverMock := driverMock.DriverDataMock{}
	driverInstance.SetInstance(driverMock)

	FindDriver(w, req)

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
