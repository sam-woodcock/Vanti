package und6iobt

import "github.com/vanti-dev/assessment-syseng-go/comm"

// Driver implements interrogation and control of a Q-SYS unD6IO-BT device.
type Driver struct {
	// Comm abstracts the underlying transport communication with the device.
	// We can assume that the underlying implementation of Comm is compatible with our device model.
	Comm comm.Transport
}
