package attributes

import "runtime"

func functionName(skip int) string {
	ret, _, _, ok := runtime.Caller(0)
	if !ok {
		return "unknown"
	}
	fn := runtime.FuncForPC(ret)
	return fn.Name()
}
