package bluetooth

// HasName defines the interface for getting a bluetooth devices announced name.
type HasName interface {
	Name() (string, error)
}
