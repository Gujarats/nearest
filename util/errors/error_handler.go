package errors

import (
	"errors"
)

func ErrorHandler(errorTag string, errorMessage string) error {
	if errorTag == "" && errorMessage == "" {
		panic("ErrorTag and ErrorMessage cannot be empty")
	}
	err := errors.New(errorTag + ":" + errorMessage)
	return err
}
