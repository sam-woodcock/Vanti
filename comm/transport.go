package comm

import "io"

// Transport defines the basic communication interface for a specific device.
// Implementations of Transport would typically wrap underlying transport protocols like TCP, UDP, or serial ports.
// Transport allows for a device driver to focus on the device protocol and not the mechanism for sending packets.
type Transport interface {
	io.ReadWriteCloser

	// Connect returns when a connection to the endpoint has been established, or none could be.
	// If the device is already connected, this method returns immediately.
	// An error is returned if a connection could not be established.
	//
	// Note that even with transport protocols that do not require a connection, like UDP,
	// Read and Write still require a Connect-ed Transport to function.
	Connect() error

	// Read implements the io.Reader interface, and reads up to len(p) bytes from an open connection.
	// Read returns an error if we are not connected to the device.
	// Read may return less than len(p) bytes if they are available rather than waiting until p is full.
	// See io.Reader for more details.
	Read(p []byte) (n int, err error)

	// Write implements the io.Writer interface, and writes p to an open connection.
	// Write returns an error if we are not connected to the device.
	// See io.Writer for more details.
	Write(p []byte) (n int, err error)

	// Close implements the io.Closer interface.
	// Close closes any open connections and interrupts any blocked Read or Write calls.
	// Subsequent calls to Close are a no-op and return the error returned from the first call, if any.
	// See io.Closer for more details.
	Close() error
}
