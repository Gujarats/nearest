package logger

import (
	"os"
	"testing"
)

func check(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

func printErrorTest(t *testing.T, actual, expected interface{}) {
	t.Errorf("Test failed expected : %s, actual : %s", expected, actual)
}

// checking if file created is exist
func TestCreateLogFile(t *testing.T) {
	_, err := CreateLogFile("/home/gujarat/testLog/", "hello.txt")
	check(t, err)

	//checking file location
	_, err = os.Stat("/home/gujarat/testLog/hello.txt")
	check(t, err)

}

func TestWriteLog(t *testing.T) {
	filePath := "/home/gujarat/testLog/"
	fileName := "hello.txt"
	inputData := "data log here"

	WriteLog(inputData, filePath, fileName)

	//checking file location
	_, err := os.Stat(filePath + fileName)
	check(t, err)

	//opening file
	f, err := os.OpenFile(filePath+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	check(t, err)

	//get the data from the file
	result := make([]byte, 100)
	_, err = f.Read(result)
	check(t, err)

	if len(result) == 0 {
		t.Error("Log Data is Empty")
	}

}
