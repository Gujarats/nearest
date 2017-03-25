package logger

import (
	"fmt"
	"log"
	"os"
	"testing"
)

var file *os.File

func init() {
	var err error
	file, err = createLogFile("./testLog/", "hello.txt")
	if err != nil {
		panic(err)
	}
}

// checking if file created is exist
func TestCreateLogFile(t *testing.T) {
	//checking file location
	_, err := os.Stat("./testLog/hello.txt")
	if err != nil {
		t.Errorf("Expecting the file exist got = %v\n", err)
	}
}

func TestWriteLog(t *testing.T) {
	filePath := "./testLog/"
	fileName := "hello.txt"
	inputData := "data log here"

	// set input error
	logger := log.New(os.Stderr,
		"Test Log :: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	logger.SetOutput(file)
	logger.Println(inputData)

	//checking file location
	_, err := os.Stat(filePath + fileName)
	if err != nil {
		t.Errorf("Error file is not exist = %v\n", err)
	}

	//opening file
	f, err := os.OpenFile(filePath+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		t.Errorf("Error opening file = %v\n", err)
	}

	//get the data from the file
	result := make([]byte, 100)
	_, err = f.Read(result)
	if err != nil {
		t.Errorf("Cannot read file = %v\n", err)
	}

	if len(result) == 0 {
		t.Error("Log Data is Empty")
	}

	//convert byte to string
	stringResult := string(result)
	fmt.Println("Result Read files = ", stringResult)
}
