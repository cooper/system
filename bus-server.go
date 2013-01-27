package system

import (
	"bufio"
	"net"
)

// listens on the specified path and returns the associated BusServer.
func BusListen(path string, bh busHandler, rh servReadHandler) (serv *BusServer, err error) {

	// resolve the address.
	addr, err := net.ResolveUnixAddr("unix", path)
	if err != nil {
		return
	}

	// begin listening.
	ln, err := net.ListenUnix("unix", addr)
	if err != nil {
		return
	}

	// no errors. create the server.
	serv = &BusServer{
		path:        path,
		listener:    ln,
		processes:   make(map[int]Process),
		busHandler:  bh,
		readHandler: rh,
		Listening:   true,
	}
	return

}

// begin the accepting loop.
func (serv *BusServer) Run() {
	for {
		conn, err := serv.listener.Accept()

		// error accepting; ignore this.
		if err != nil {
			continue
		}

		// begin the connection goroutine.
		go serv.handleConnection(conn)

	}
}

// handle incoming data from a connection.
func (serv *BusServer) handleConnection(conn net.Conn) {

	// create a buffered reader.
	reader := bufio.NewReader(conn)

	// attempt to read a line.
	for {
		line, err := reader.ReadString('\n')

		// error occured before reaching delimiter.
		if err != nil {
			return
		}

		// no error. handle the line.
		serv.readHandler(serv, conn, line)

	}
}
