package system

/*
   this file implements a system logger.
   initially, the process starts with a logger to STDOUT. Typically, no logging is done
   before the logger switches over to a system logging bus.
*/

import (
	"fmt"
	"log"
	"os"
)

// the main system logger instance.
// this is setup when Register() is called. (see: interface.go)
var Logger *log.Logger

// creates the initial logger.
func createLogger(name string) (logger *log.Logger) {
	logger = log.New(os.Stdout, fmt.Sprintf("[%s] ", name), log.Ldate|log.Ltime|log.Lshortfile)
	return
}
