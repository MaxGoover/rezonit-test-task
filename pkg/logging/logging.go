package logging

import (
	"log"
	"os"
)

var Info *log.Logger
var Error *log.Logger

func init() {
	Info = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	Error = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
}
