package system

import (
	"bufio"
	"net"
)

// master event handler function type.
type busHandler func(source Process, command string, params map[string]interface{})

// single event hander function type.
type eventHandler func(source Process, params map[string]interface{})

// represents a connection to a bus.
type BusConnection struct {
	path     string        // connect address
	socket   net.Conn      // there's no reason to restrict to Unix sockets
	incoming *bufio.Reader // data reader - currently used for clients only
	handler  busHandler    // the function called when an event is received
}
