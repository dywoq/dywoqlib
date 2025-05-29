package container

// ActualLengthGetter defines a method for getting the actual length.
type ActualLengthGetter interface {
	ActualLength() int
}

// InitialLengthGetter defines a method for getting the initial length.
// Usually it's implemented by Fixed structure.
type InitialLengthGetter interface {
	InitialLength() int
}
