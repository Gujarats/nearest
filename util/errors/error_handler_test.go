package errors

import (
	"errors"
	"testing"
)

func TestErrorHandler(testing *testing.T) {
	var testCases = []struct {
		errorTag     string
		errorMessage string
		expected     error
	}{
		{
			"Hello",
			"This is Error Message for you darling",
			errors.New("Hello" + ":" + "This is Error Message for you darling"),
		},

		{
			"Hi",
			"How about we go to watch movie tonight",
			errors.New("Hi" + ":" + "How about we go to watch movie tonight"),
		},
	}

	for _, t := range testCases {
		actual := ErrorHandler(t.errorTag, t.errorMessage)
		if actual.Error() != t.expected.Error() {
			testing.Errorf("Test failed expected : %s, actual : %s", t.expected, actual)
		}
	}
}
