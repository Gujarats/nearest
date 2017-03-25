package logger

import "os"

// NOTE : becareful to use this function. if you got permission denied then it will loop forever.
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
