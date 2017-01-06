package costumeLog

import (
	"log"
	"os"
)

var (
	Trace *log.Logger
)

func Init() {

	Trace = log.New(os.Stderr,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

}
