package system

/*
   this file implements a system logger.
   initially, the process starts with a logger to STDOUT. Typically, no logging is done
   before the logger switches over to a system logging bus.
*/

import (
	"log"
	"os"
)

// the logwriter type is used primarily as an io.Writer for the log class.
type logwriter bool

// implements io.Writer
func (l logwriter) Write(p []byte) (n int, err error) {

	// if the bus is connected, send a "log" message.
	if LogBusConn.Connected {
		LogBusConn.Send("log", map[string]interface{}{
			"message": p,
		})
	}
	n = len(p)
	return

	// otherwise, write to STDOUT.
	n, err = os.Stdout.Write(p)
	return

}

// the main system logger instance.
// this is setup when Register() is called. (see: interface.go)
var Logger *log.Logger

// creates the initial logger.
func createLogger(name string) (logger *log.Logger) {
	var writer logwriter = true
	logger = log.New(writer, "["+name+"] ", log.Ldate|log.Ltime|log.Lshortfile)
	return
}

// called after connecting to log bus.
func runLogger() {

	// now that we have connected, send greeting.
	LogBusConn.Send("register", map[string]interface{}{
		"programName": Self.name,
	})

	// replace the current logger with one that writes to the logging bus.

	// begin the loop.
	go LogBusConn.Run()
}
