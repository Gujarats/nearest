package logger

import (
	"log"
	"os"
)

var Trace *log.Logger
var FileLog *os.File

func InitLogger(prefix, filePath, fileName string) {
	SetPrefix(prefix)
	CreateLogFile(filePath, fileName)
}

func SetPrefix(prefix string) {
	Trace = log.New(os.Stderr,
		prefix,
		log.Ldate|log.Ltime|log.Lshortfile)

}

func CreateLogFile(filePath, fileName string) (*os.File, error) {
	SetPrefix("FILE LOG : ")
	f, err := os.OpenFile(filePath+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		// error file is not created
		err = os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			Trace.Println(err)
			return nil, err
		}
		CreateLogFile(filePath, fileName)
	}

	return f, nil
}

func WriteLog(dataLog, filePath, fileName string) {
	//assign to global variable
	f, err := CreateLogFile(filePath, fileName)
	if err != nil {
		Trace.Println(err)
		return
	}

	Trace.SetOutput(f)
	Trace.Println(dataLog, err)
}
