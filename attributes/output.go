package attributes

import (
	"fmt"
	"log"
)

func output(message string, mode Mode) {
	if mode == StrictMode {
		log.Fatal(message)
	}
	fmt.Println(message)
}
