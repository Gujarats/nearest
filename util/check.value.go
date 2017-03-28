package util

import "reflect"

// Check if all input values in not nil.
func CheckValue(values ...string) bool {
	for _, value := range values {
		if value == "" {
			return false
		}
	}

	return true
}

// passing struct and check all the attribute,
// false if there is empty value.
func CheckAttribute(input interface{}) bool {
	object := reflect.ValueOf(input)

	for index := 0; index < object.NumField(); index++ {

		if IsZeroOfUnderlyingType(object.Field(index).Interface()) {
			return false
		}
	}

	return true
}

// compare the object value to non assign value.
func IsZeroOfUnderlyingType(objectValue interface{}) bool {
	return reflect.DeepEqual(objectValue, reflect.Zero(reflect.TypeOf(objectValue)).Interface())
}
