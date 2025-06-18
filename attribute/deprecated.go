package attribute

import (
	"fmt"
	"strings"
)

// Deprecated is an attribute function that generates warning about deprecated function.
// DOES NOT automatically returns zero values (e.g., "", 0, nil etc.).
func Deprecated() {
	m := management{}
	targetSkip, sourceSkip := m.skipNums()
	target := m.funcInfo(targetSkip)
	source := m.funcInfo(sourceSkip)

	strs := make([]string, 5)
	strs = append(strs, "attribute.Deprecated: ", target, " is deprecated; ")
	strs = append(strs, "the source of the warning: ", source)

	res := strings.Join(strs, "")
	fmt.Println(res)
}
