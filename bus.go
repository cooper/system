package system

// master event handler function type.
type busHandler func(command string, params map[string]interface{})

// single event hander function type.
type eventHandler func(params map[string]interface{})

// represents a connection to a bus.
type BusConnection struct {
	path     string   // connect address
	socket   net.Conn // there's no reason to restrict to Unix sockets
	incoming *bufio.Reader
	handler  busHandler
}

