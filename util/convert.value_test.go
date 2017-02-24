package util

import (
	"fmt"
	"testing"
)

func TestConverToFloat64(t *testing.T) {
	type testObject struct {
		Input1 string
		Input2 string
		Input3 string
		Input4 string
	}

	testObjects := []testObject{
		{Input1: "1.2", Input2: "2.43434", Input3: "5.23232", Input4: "8.3423432"},
		{Input1: "2.2", Input2: "23.43434", Input3: "15.23232", Input4: "18.3423432"},
	}

	for _, test := range testObjects {
		result := ConvertToFloat64(test.Input1, test.Input2, test.Input3, test.Input3)
		fmt.Printf(" result = %+v\n", result)
	}

}
