package attribute

import (
	"fmt"
	"strings"
)

// Deprecated is an attribute-function that generates warning about deprecated function.
// DOES NOT automatically returns zero values (e.g., "", 0, nil etc.).
//
// If custom event is not nil (it's set by SetEvent), 
// then it uses the custom event instead of outputting the warning.
func Deprecated() {
	m := management{}
	targetSkip, sourceSkip := m.skipNums()
	target := m.funcInfo(targetSkip)
	source := m.funcInfo(sourceSkip)

	strs := make([]string, 5)
	strs = append(strs, "attribute.Deprecated: ", target, " is deprecated; ")
	strs = append(strs, "the source of the warning: ", source)

	res := strings.Join(strs, "")
	if event != nil {
		event()
		return
	}
	fmt.Println(res)
}
