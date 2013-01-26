package system

/* this file contains the programming interface used in client applications. */

// the SystemBus type defines the object system.Bus.
type SystemBus struct {
}

// Bus.RegisterListener(command, handler)
func (bus *SystemBus) RegisterListener(command string, handler eventHandler) {

}

// registers the process to the process manager and system bus.
func Register(name string, version string, description string) {

}
