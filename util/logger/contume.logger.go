package logger

import (
	"log"
	"os"
)

func createLogFile(filePath, fileName string) (*os.File, error) {
	f, err := os.OpenFile(filePath+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		// error file is not created
		err = os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			panic(err)
			return nil, err
		}
		createLogFile(filePath, fileName)
	}

	return f, nil
}

// use this PrintLog use print the error on the terminal without exitting apps.
func PrintLog(logger *log.Logger, file *os.File, errorLog string) {

	// save the log out put to file
	logger.SetOutput(file)
	logger.Println(errorLog)
}

// use this Fatal log to exit the application whenever got error.
func FatalLog(logger *log.Logger, file *os.File, errorLog string) {
	// save the log out put to file
	logger.SetOutput(file)
	logger.Fatalln(errorLog)
}
