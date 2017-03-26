package logger

import (
	"fmt"
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

	// remove the file first if exist
	_, err := os.Stat(filePath + fileName)
	if err != nil {
		// file is not exist, then continue to the test
		return
	}

	os.Remove(filePath + fileName)

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

func TestCheckfile(t *testing.T) {
	file, err := os.OpenFile(filePath+fileName, os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	result := make([]byte, 100)
	file.Read(result)

	fmt.Println(result)
}
