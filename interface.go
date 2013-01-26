package system

/* this file contains the programming interface used in client applications. */

import "os"

var (
	Self           Process
	ProcessBusConn *BusConnection
	LogBusConn     *BusConnection
)

// registers the process to the process manager and system bus.
// Register("Some program", "1.0", "a sample program")
func Register(name string, version string, description string) {

	// create the initial logger.
	Logger = createLogger(name)

	// create this process's Process.
	Self = newClientProcess(os.Getpid())

	// connect to the system bus.
	ProcessBusConn, err := BusConnect("/System/Bus/processbus", clientHandler, jsonDataHandler)
	if err != nil {
		// die...
	}

	// connect to the logging bus.
	LogBusConn, err = BusConnect("/System/Bus/logbus", nil, nil)

	// run the loops.
	go conn.Run()
	go logconn.Run()

}

/*##########################
### SYSTEM BUS INTERFACE ###
##########################*/

func init() {
	listeners = make(map[string]eventHandler)
}

var listeners map[string]eventHandler

// the SystemBus type defines the object system.Bus.
type SystemBus struct {
}

// registers an event listener.
// Bus.RegisterListener(command, handler)
func (bus *SystemBus) RegisterListener(command string, handler eventHandler) {

	// store the event handler.
	listeners[command] = handler

}

// internal event handler. (type busHandler)
func clientHandler(source Process, command string, params map[string]interface{}) {
	if listeners[command] != nil {
		listeners[command](source, params)
	}
}
