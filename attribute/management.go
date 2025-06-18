package attribute

import (
	"runtime"
)

type management struct{}

func (management) funcInfo(skip int) (string) {
	ret, _, _, ok := runtime.Caller(skip)
	if !ok {
		return "unknown"
	}
	fn := runtime.FuncForPC(ret)
	name := fn.Name()
	return name
}

func (management) skipNums() (int, int) {
	// 2 - skip num of caller of an attributes
	// 3 - skip num of source of the warning
	return 2, 3
}

