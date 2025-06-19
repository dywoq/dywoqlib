package attribute

var event func()

// SetEvent sets custom event to attributes-functions.
func SetEvent(value func()) {
	event = value
}

// ResetEvent sets the custom event to nil.
func ResetEvent() {
	event = nil
}
