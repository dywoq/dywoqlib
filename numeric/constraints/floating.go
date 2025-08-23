package constraints

// Floating is a constraint of float types such as float32, float64.
type Floating interface {
	~float32 | ~float64
}
