package system

/* system path constants */

const (
	PATH_RAMDISK = ""
	PATH_SYSDISK = "/drive" //                        // main drive mountpoint (/Volumes/System)

	PATH_USER        = "/Files"        //             // user files spun across all drives
	PATH_VOLUME      = "/Volumes"      //             // mounted volumes
	PATH_APPLICATION = "/Applications" //             // application bundles

	PATH_SYSTEM     = "/System"                       // files belonging to the system itself
	PATH_SYSLIBRARY = "/System/Library"               // modules, libraries, extensions, etc.
	PATH_SYSLIBSO   = "/System/Library/SharedObjects" // dynamically loaded libraries (i.e. /lib)
	PATH_SYSBUS     = "/System/Bus"                   // system bus sockets
	PATH_SYSDEVICE  = "/System/Devices"               // Linux device nodes (eq. /dev)
	PATH_SYSLOG     = "/System/Logs"                  // generated logs
	PATH_SYSPROCESS = "/System/Processes"             // process files (similar to /proc)
)
