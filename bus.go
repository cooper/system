package system

import (
	"bufio"
	"net"
)

// data read handler.
type readHandler func(conn *BusConnection, data []byte) bool
type servReadHandler func(serv *BusServer, conn net.Conn, line string) bool

// master event handler function type.
type busHandler func(source Process, command string, params map[string]interface{})

// single event hander function type.
type eventHandler func(source Process, params map[string]interface{})

// represents a connection to a bus. this is used in clients.
type BusConnection struct {
	path        string        // connect address
	socket      net.Conn      // there's no reason to restrict to Unix sockets
	incoming    *bufio.Reader // data reader - currently used for clients only
	outgoing    *bufio.Writer // data writer - currently used for clients only
	busHandler  busHandler    // the function called when an event is received
	readHandler readHandler   // the function called when data is received
	Connected   bool          // true if the bus is currently connected
}

// represents a bus server. this is used in servers.
type BusServer struct {
	path        string          // the listen address
	listener    net.Listener    // typically a *UnixListener
	processes   map[int]Process // the currently connected processes
	busHandler  busHandler      // the function called when an event is received
	readHandler servReadHandler // the function called when data is received
	Listening   bool            // true if the bus is currently listening
}
