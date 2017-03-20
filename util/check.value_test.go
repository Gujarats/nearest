package util

import (
	"testing"

	"github.com/Gujarats/API-Golang/model/driver"
)

func TestCheckValue(t *testing.T) {
	type testObject struct {
		Input    string
		Input2   string
		Input3   string
		Expected bool
	}

	testObjects := []testObject{
		{"hello", "asdf", "test", true},
		{"", "asdf", "", false},
		{"cek", "asdf", "", false},
	}

	for _, test := range testObjects {
		actual := CheckValue(test.Input, test.Input2, test.Input3)
		if test.Expected != actual {
			t.Errorf("Error :: actual = %v, expected = %v", actual, test.Expected)
		}
	}
}

func TestCheckAttributeTrue(t *testing.T) {
	// testing model driver
	driverData := driver.DriverData{
		Id:     "asdf",
		Name:   "name",
		Status: true,
		Location: driver.GeoJson{
			Type:        "test",
			Coordinates: []float64{2.34, 3.34},
		},
	}

	expected := true

	actual := CheckAttribute(driverData)

	if actual != expected {
		t.Errorf("Error :: actual = %v, expected = %v", actual, expected)
	}

}

func TestCheckAttributeFalse(t *testing.T) {
	// testing model driver
	driverData := driver.DriverData{
		Name: "name",
	}

	expected := false

	actual := CheckAttribute(driverData)

	if actual != expected {
		t.Errorf("Error :: actual = %v, expected = %v", actual, expected)
	}

}

func TestCheckAttributeIntTrue(t *testing.T) {
	//create struct that has int value
	type TestObject struct {
		Value  int
		Value2 float64
	}

	testObject := TestObject{
		Value:  1,
		Value2: 2.34,
	}

	expected := true

	actual := CheckAttribute(testObject)

	if actual != expected {
		t.Errorf("Error :: actual = %v, expected = %v", actual, expected)
	}

}

func TestCheckAttributeIntFalse(t *testing.T) {
	//create struct that has int value
	type TestObject struct {
		Value  int
		Value2 float64
	}

	testObject := TestObject{
		Value: 1,
	}

	expected := false

	actual := CheckAttribute(testObject)

	if actual != expected {
		t.Errorf("Error :: actual = %v, expected = %v", actual, expected)
	}

}
