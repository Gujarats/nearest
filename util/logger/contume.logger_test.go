package logger

import (
	"fmt"
	"log"
	"os"
	"testing"
)

var file *os.File
var logger *log.Logger

func init() {
	var err error
	file, err = createLogFile("./testLog/", "hello.txt")
	if err != nil {
		panic(err)
	}
	logger = log.New(os.Stderr,
		"Test Log :: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

// checking if file created is exist
func TestCreateLogFile(t *testing.T) {
	//checking file location
	_, err := os.Stat("./testLog/hello.txt")
	check(t, err)
}

func TestWriteLog(t *testing.T) {
	filePath := "./testLog/"
	fileName := "hello.txt"
	inputData := "data log here"
	PrintLog(logger, file, inputData)

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

	//convert byte to string
	stringResult := string(result)
	fmt.Println("Result Read files = ", stringResult)
}

// ============ PRIVATE FUNCTION ============//

// checking if erro not nil.
func check(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

//pritnt the error if the result is not equal with the expected result.
func printErrorTest(t *testing.T, actual, expected interface{}) {
	t.Errorf("Test failed expected : %s, actual : %s", expected, actual)
}
