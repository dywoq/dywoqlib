package attribute

import (
	"fmt"
	"log"
	"runtime"
)

type management struct{}

func (m management) output(message string, mode Mode) {
	switch mode {
	case SoftMode:
		fmt.Println(message)
	case StrictMode:
		log.SetFlags(0)
		log.Fatalln(message)
	}
}

func (m management) functionName(skip int) string {
	zero := "unknown" // the variable is to be returned if error happens
	ret, _, _, ok := runtime.Caller(skip)
	if !ok {
		return zero
	}
	fn := runtime.FuncForPC(ret)
	return fn.Name()
}

func (m management) targetNumberSkip() int {
	return 2
}

func (m management) sourceNumberSkip() int {
	return 3
}
