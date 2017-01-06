package costumeLog

import (
	"log"
	"os"
)

var Trace *log.Logger

func SetPrefix(prefix string) {

	Trace = log.New(os.Stderr,
		prefix,
		log.Ldate|log.Ltime|log.Lshortfile)

}
