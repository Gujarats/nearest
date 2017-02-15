package review

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/training_project/model/review"
	"github.com/training_project/model/review/instance"
	"github.com/training_project/model/review/mock"
)

func TestEndPointSatusOk(t *testing.T) {
	data := url.Values{}
	data.Set("shop_id", "17112321") // dummy shopID value. inside database is not exist

	//we pass a dummy value to pass the required params
	req := httptest.NewRequest("POST", "/", bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	w := httptest.NewRecorder()

	//  create mock object to test handler. Must called first before handler
	// mock data is exist
	reviewMock := reviewMock.ReviewDataMock{IsDataExist: true}
	reviewInstance.SetInstance(reviewMock)

	// we're going to test handler
	CheckDataExist(w, req)

	// check the status code
	if w.Code != http.StatusOK {
		t.Errorf("Error Code !! actual : %d, expected : %d\n", w.Code, http.StatusOK)
	}

	// check the response
	resp := w.Body.Bytes()
	if resp == nil {
		t.Error("Error Body result Empty")
	}

	reviewResp := review.ReviewResponse{}
	err := json.Unmarshal(resp, &reviewResp)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("result response = %+v\n", reviewResp)
	if reviewResp.Message != "Data Exist" {
		t.Errorf("Error response:: actual = %v, expected = %v\n", reviewMock.Message, "Data Exist")
	}

}

func TEstEndPointStatusBadReq(t *testing.T) {
	// create data body
	data := url.Values{}
	data.Add("shop_id", "12342")

	// create request
	req := httptest.NewRequest("POST", "/", bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp := httptest.NewRecorder()

	reviewMock := reviewMock.ReviewDataMock{IsDataExist: false}
	reviewInstance.SetInstance(reviewMock)

	CheckDataExist(resp, req)

	if resp.Code != http.StatusBadRequest {
		t.Errorf("Error Code !! actual : %d, expected : %d\n", resp.Code, http.StatusBadRequest)
	}

	// check the response
	respData := resp.Body.Bytes()
	if respData == nil {
		t.Error("Error Body result Empty")
	}

	reviewResponse := review.ReviewResponse{}
	err := json.Unmarshal(respData, &reviewResponse)
	if err != nil {
		t.Error(err)
	}

	if reviewResponse.Message != "Data is Not Exist" {
		t.Errorf("Error response:: actual = %v, expected = %v\n", reviewResponse.Message, "Data Exist")
	}
}
