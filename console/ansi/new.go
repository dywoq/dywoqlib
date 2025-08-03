package ansi

// New creates a new ANSI message. The colors are None automatically.
func New(value string) Base {
	return &message{None, None, value}
}
