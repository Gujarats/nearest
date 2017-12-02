package util

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

// Test if coverting to radian is more than zero
func TestToRadians(t *testing.T) {
	testObjects := []struct {
		Input    float64
		Expected float64
	}{
		{1.2, 0.0},
		{1.4, 0.0},
		{0.1, 0.0},
		{0.0, 0.0},
	}

	for _, test := range testObjects {
		actual := toRadians(test.Input)
		if actual < test.Expected {
			t.Errorf("Input = %v, Error actual = %v , expected = %v\n", test.Input, actual, test.Expected)
		}
	}
}

// Test Distance see if tow point distance calculation is more than zero
func TestDistance(t *testing.T) {
	testObjects := []struct {
		Lat1     float64
		Lon1     float64
		Lat2     float64
		Lon2     float64
		Expected float64
	}{
		{Lat1: -6.8915208, Lon1: 107.6100268, Lat2: 6.8937359, Lon2: 107.6083563, Expected: 0.0},
	}

	for _, test := range testObjects {
		actual := Distance(test.Lat1, test.Lon1, test.Lat2, test.Lon2)

		if actual < test.Expected {
			t.Errorf("Error actual = %v, Expect more then zero = %v\n", actual, test.Expected)
		}
	}
}

func TestGetNearestLocation(t *testing.T) {
	testObjects := []struct {
		myLoc     Location
		othersLoc []Location
		expected  Location
	}{
		{
			myLoc: Location{Lat: -6.8836631, Lon: 107.5969201},
			othersLoc: []Location{
				// cileunyi
				{
					Lat: -6.9271271,
					Lon: 107.7190409,
				},

				// kampung tulip
				{

					Lat: -6.963114,
					Lon: 107.6612395,
				},
			},
			expected: Location{
				Lat: -6.963114,
				Lon: 107.6612395,
			},
		},
	}

	for _, testObject := range testObjects {
		actual := GetNearestLocation(testObject.myLoc, testObject.othersLoc)
		assert.Equal(t, actual.RLocation, testObject.expected)
	}
}
