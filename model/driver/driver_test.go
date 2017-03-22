package driver

import "testing"

func TestGetDistrictFormat(t *testing.T) {
	testObjects := []struct {
		City       string
		IdDistrict string
		Expected   string
	}{
		{City: "testCity", IdDistrict: "testIdDistrict", Expected: "testCity_district_testIdDistrict"},
		{City: "testCity2", IdDistrict: "testIdDistrict2", Expected: "testCity2_district_testIdDistrict2"},
	}

	for _, testObject := range testObjects {
		actual := getFormatDistrict(testObject.City, testObject.IdDistrict)
		if actual != testObject.Expected {
			t.Errorf("Error actual = %v, expected %v\n", actual, testObject.Expected)
		}
	}
}
