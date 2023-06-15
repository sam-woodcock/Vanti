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
	ConnectionConnectedAVRCPNotSupported
	ConnectionConnectedAVRCPSupported
	ConnectionConnectedAVRCPPDUSupported
)

func GetConnectionStatus(connection int) Connection {
	switch Connection(connection) {
	case ConnectionUnknown:
		return ConnectionUnknown
	case ConnectionNotConnected:
		return ConnectionNotConnected
	case ConnectionConnected, ConnectionConnectedAVRCPNotSupported, ConnectionConnectedAVRCPSupported, ConnectionConnectedAVRCPPDUSupported:
		return ConnectionConnected
	default:
		return ConnectionUnknown
	}
}

// 0 = Idle
// 1 = Discoverable
// 2 = Connected – Unknown AVRCP support
// 3 = Connected – AVRCP Not Supported
// 4 = Connected – AVRCP Supported
// 5 = Connected – AVRCP & PDU Supported
