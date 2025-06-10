package attribute

import "runtime"

func functionName(skip int) string {
	ret, _, _, ok := runtime.Caller(skip)
	if !ok {
		return "unknown"
	}
	fn := runtime.FuncForPC(ret)
	return fn.Name()
}
