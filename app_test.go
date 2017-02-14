package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/training_project/controller/review"
	"github.com/training_project/controller/review/struct"
	"github.com/training_project/model/review/interface"
)

func TestEndPoint(t *testing.T) {
	data := url.Values{}
	data.Set("shop_id", "171")

	req := httptest.NewRequest("POST", "/", bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	w := httptest.NewRecorder()

	//  create mock object to test handler
	reviewMock := reviewMethod.ReviewDataMock{ShopID: 12}
	reviewStruct.SetStruct(reviewMock)

	// we're going to test handler
	review.CheckDataExist(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Error Code !! actual : %d, expected : %d\n", w.Code, http.StatusOK)
	}
}
