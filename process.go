package system

import "os"

/*#########################
### SystemProcess class ###
#########################*/

// Process objects typically inherit from the SystemProcess class, which itself inherits
// from Go's os.Process.
type SystemProcess struct {
	*os.Process
}

// returns a new SystemPrcoess
func newSystemProcess(pid int) *SystemProcess {
	proc, _ := os.FindProcess(pid)
	return &SystemProcess{proc}
}

// returns the numerical process ID
func (p *SystemProcess) PID() int {
	return p.Pid // provided by os.Process
}

/* [Process interface specification] */

// this defines the methods that must be available to system process objects.
type Process interface {

	// sends a message via the system communication bus.
	Send(message string, data map[string]interface{})

	/* Provided by SystemProcess */

	// returns the numerical process ID
	PID() int
}

/*#########################
### ClientProcess class ###
#########################*/
// this class complies with Process interface and is
// intended for use in everyday programs.

type ClientProcess struct {
	*SystemProcess
}

// returns a new ClientProcess.
// this is a low level interface. high-level interface will be provided by FindProcess().
func newClientProcess(pid int) *ClientProcess {
	return &ClientProcess{newSystemProcess(pid)}
}

// sends a message directly to a process.
func (p *ClientProcess) Send(message string, data map[string]interface{}) {
}
