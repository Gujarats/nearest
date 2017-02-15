package logger

import (
	"log"
	"os"
)

var trace *log.Logger
var filePathVar string
var fileNameVar string
var prefixVar string

func InitLogger(prefix, filePath, fileName string) {
	setPrefix(prefix)
	createLogFile(prefix, filePath, fileName)
	filePathVar = filePath
	fileNameVar = fileName
	prefixVar = prefix
}

func setPrefix(prefix string) {
	trace = log.New(os.Stderr,
		prefix,
		log.Ldate|log.Ltime|log.Lshortfile)

}

func createLogFile(prefix, filePath, fileName string) (*os.File, error) {
	setPrefix(prefix)
	f, err := os.OpenFile(filePath+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		// error file is not created
		err = os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			trace.Println(err)
			return nil, err
		}
		createLogFile(prefix, filePath, fileName)
	}

	return f, nil
}

// use this PrintLog use print the error on the terminal without exitting apps.
func PrintLog(errorLog string) {
	//assign to global variable
	if filePathVar == "" || fileNameVar == "" || prefixVar == "" {
		trace.Fatalln("Error file path and Name empty prefix Cannot be Empty")
	}

	f, err := createLogFile(prefixVar, filePathVar, fileNameVar)
	if err != nil {
		trace.Fatalln(err)
		return
	}
	trace.Println(errorLog)
	// save the log out put to file
	trace.SetOutput(f)
	trace.Println(errorLog, err)
}

// use this Fatal log to exit the application whenever got error.
func FatalLog(errorLog string) {
	//assign to global variable
	if filePathVar == "" || fileNameVar == "" || prefixVar == "" {
		trace.Fatalln("Error file path and Name empty")
	}

	f, err := createLogFile(prefixVar, filePathVar, fileNameVar)
	if err != nil {
		trace.Fatalln(err)
		return
	}
	trace.Println(errorLog)
	// save the log out put to file
	trace.SetOutput(f)
	trace.Fatalln(errorLog, err)
}
