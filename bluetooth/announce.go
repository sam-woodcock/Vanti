package bluetooth

// HasAnnounce provides functionality for triggering this device to activate the pairing flow.
// After calling Announce, this device should be discoverable by Bluetooth sources to connect to using their
// native Bluetooth pairing interfaces.
type HasAnnounce interface {
	Announce() error
}
