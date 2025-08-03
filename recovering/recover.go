package recovering

import (
	"log"
)

// Recover catches the encountered panic and returns it. Unlike builtin recover(),
// the function outputs the log message about panic.
// The message of caught panic would look like this:
// recovering.Recover(): caught the panic: <caught panic() value>
func Recover() (r any) {
	if r = recover(); r != nil {
		log.Printf("recovering.Recover(): caught the panic: %v", r)
	}
	return
}
