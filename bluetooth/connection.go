package bluetooth

// HasConnection returns the connection status for this device, i.e. whether a phone is connected.
// ConnectionChanged blocks until the connection status is different from last.
type HasConnection interface {
	ConnectionChanged(last Connection) (Connection, error)
}

// Connection indicates whether a remote device is connected via Bluetooth to a device.
type Connection int

const (
	// ConnectionUnknown is used to indicate that we don't know the connection status.
	// It should be used as a response under error conditions and can be used as a parameter for ConnectionChanged.
	ConnectionUnknown Connection = iota
	// ConnectionNotConnected indicates that no Bluetooth connection is active.
	ConnectionNotConnected
	// ConnectionConnected indicates that there is an active Bluetooth connection.
	ConnectionConnected
)
