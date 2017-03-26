package logger

import (
	"log"
	"os"
	"testing"
)

var logger *log.Logger
var filePath string
var fileName string

func init() {
	filePath = "./testLog/"
	fileName = "hello.txt"
	logger = log.New(os.Stderr,
		"Test Log :: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

// checking if file created is exist
func TestCreateLogFile(t *testing.T) {
	//checking file location
	_, err := os.Stat(filePath + fileName)
	if err != nil {
		t.Errorf("Expecting the file exist got = %v\n", err)
	}
}

func TestWriteLog(t *testing.T) {
	// set file and print some error to the error message to file.
	inputData := "data log here"

	file, err := createLogFile(filePath, fileName)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	logger.SetOutput(file)
	logger.Println(inputData)
}
