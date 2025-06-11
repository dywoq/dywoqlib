package attribute

// Mode indicates the mode of an attribute warning message.
// Its enumerations are SoftMode, StrictMode.
type Mode int

const (
	// SoftMode is a mode that doesn't allows warning to terminate the program.
	SoftMode Mode = iota

	// StrictMode is a mode which makes warning terminate the program.
	StrictMode
)
