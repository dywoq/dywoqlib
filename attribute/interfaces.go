package attribute

// Attributable is an interface with method Name().
type Attributable interface {
	// Name returns the name of an attribute.
	// The name should be formatted this way:
	// `package_name.Attribute`
	Name() string
}

// Callable is an interface with method Call().
type Callable interface {
	// Call activates the warning and outputs it into console.
	// If mode is set to SoftMode, program will be not executed.
	// Otherwise, if it's set to StrictMode, program will be immediately executed.
	// The message should contain:
	//
	// - name of an attribute;
	//
	// - the function where Call() was used;
	//
	// - the source of warning.
	Call(mode Mode)
}
