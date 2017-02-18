package util

import "testing"

func TestCheckValueTrue(t *testing.T) {
	expected := true
	actual := CheckValue("hello", "asdf", "test")

	if actual != expected {
		t.Errorf("Error :: actual = %v, expected = %v", actual, expected)
	}
}

func TestCheckValueFalse(t *testing.T) {
	expected := false
	actual := CheckValue("hello", "", "test")

	if actual != expected {
		t.Errorf("Error :: actual = %v, expected = %v", actual, expected)
	}
}
