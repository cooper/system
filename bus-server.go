package system

import "net"

// listens on the specified path and returns the associated BusServer.
func BusListen(path string, bh busHandler, rh readHandler) (serv *BusServer, err error) {

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
		path:      path,
		listener:  ln,
		processes: make(map[int]Process),
	}
	return

}
