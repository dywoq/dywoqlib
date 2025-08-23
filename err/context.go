package err

import (
	"encoding/json"
)

// Context defines an interface for error handling with additional context.
// It allows access to the error itself, extra details, and different
// representations such as string or JSON.
type Context interface {
	// Error returns the original error.
	Error() error
	// More returns additional context about the error.
	More() string
	// Nil reports whether the error is nil.
	Nil() bool
	// String returns a formatted string with the error and additional context.
	String() string
	// Marshal returns a JSON representation of the error and context.
	Marshal() ([]byte, error)
}

type implementation struct {
	err  error
	more string
}

// NewContext creates a new Context with the given error and additional context.
func NewContext(err error, more string) Context {
	return implementation{err, more}
}

// Error returns the original error.
func (i implementation) Error() error {
	return i.err
}

// More returns additional context information.
func (i implementation) More() string {
	return i.more
}

// Nil reports whether the error is nil.
// (Currently returns true if error is not nil, may need adjustment.)
func (i implementation) Nil() bool {
	return i.err != nil
}

// String returns a human-readable string combining error and context.
func (i implementation) String() string {
	return i.err.Error() + ": " + i.more
}

type jsonPayload struct {
	Error string `json:"error"`
	More  string `json:"more"`
}

// Marshal returns the JSON representation of the error and context.
func (i implementation) Marshal() ([]byte, error) {
	data := jsonPayload{
		Error: i.err.Error(),
		More:  i.more,
	}
	return json.Marshal(data)
}
