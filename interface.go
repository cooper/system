package system

/* this file contains the programming interface used in client applications. */

// registers the process to the process manager and system bus.
// Register("Some program", "1.0", "a sample program")
func Register(name string, version string, description string) {

	// create the initial logger.
	Logger = createLogger(name)

	// connect to the system bus.
	conn, err := BusConnect("/System/Bus/processbus", clientHandler)
	if err != nil {
		// die...
	}

	// run the loop.
	go conn.Run()

}

/*##########################
### SYSTEM BUS INTERFACE ###
##########################*/

var listeners map[string]eventHandler

// the SystemBus type defines the object system.Bus.
type SystemBus struct {
}

// registers an event listener.
// Bus.RegisterListener(command, handler)
func (bus *SystemBus) RegisterListener(command string, handler eventHandler) {

	// initialize listeners.
	if listeners == nil {
		listeners = make(map[string]eventHandler)
	}

	// store the event handler.
	listeners[command] = handler

}

// internal event handler. (type busHandler)
func clientHandler(source Process, command string, params map[string]interface{}) {
	if listeners[command] != nil {
		listeners[command](source, params)
	}
}
