package recovering

import (
	"log"
)

// Recover catches the encountered panic.
// If event is not nil, it will be used instead of default logging.
// If cleanup is not nil, cleanup will be ran after the logging or the custom event finished.
func Recover(event, cleanup func()) (r any) {
	if r = recover(); r != nil {
		if event == nil {
			log.Printf("recovering.Recover(): caught the panic: %v", r)
		} else {
			event()
		}
		if cleanup != nil {
			cleanup()
		}
	}
	return
}
