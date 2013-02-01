# system

This is a system library that provides a message bus; process, user, and device
management; and much more.

## Goals of this software

* __initial daemon__: process launcher and init daemon. (LaunchManager)
* __communication__: interprocess communication mechanism. (CommunicationManager)
* __logging__: system logging bus. (LogManager)
* __users__: user, files, and password management. (UserManager)
* __devices__: device management service. (DeviceManager)

## Browing the source

If you are unfamiliar with Go package source files, a simple explanation is that the
file location of source code typically does not matter. Similarly, the order in which
functions, constants, types, variables, etc. are defined does not matter. Any of these
could be moved from file to file without changing the functionality of the program.
  
However, this projects organizes related code into files, as seen below:

* __`interface.go`__: contains most of the exported functions for use in everyday programs.
* __`bus.go`__: includes code used by files involving system buses.
* __`bus-server.go`__: provides a simple listening mechanism for listening system bus servers.
* __`bus-client.go`__: provides a simple means of connecting to a system bus server.
* __`constant.go`__: contains exported system constants.
* __`logger.go`__: provides system.Logger, an instance of Go's Logger connected to a system log bus.
* __`process.go`__: includes the Process type specification and other code related to the objective system process programming interface.

## System-provided constants

These constants are provided for use throughout the system. Because they are part of the
'system' package, they must be accessed correctly, i.e.: `system.PATH_SYSLIBSO`.

```go
/* Path constants */
const (
	PATH_RAMDISK = ""
	PATH_SYSDISK = "/drive" // Volumes/System

	PATH_USER        = "/Files"
	PATH_VOLUME      = "/Volumes"
	PATH_APPLICATION = "/Applications"

	PATH_SYSTEM     = "/System"
	PATH_SYSLIBRARY = "/System/Library"
	PATH_SYSLIBSO   = "/System/Library/SharedObjects"
	PATH_SYSBUS     = "/System/Bus"
	PATH_SYSDEVICE  = "/System/Devices"
	PATH_SYSLOG     = "/System/Logs"
	PATH_SYSPROCESS = "/System/Processes"
)
```
