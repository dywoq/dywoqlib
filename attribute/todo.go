package attribute

import (
	"fmt"
	"strings"
)

// Todo is an attribute function that generates warning about unimplemented function.
// DOES NOT automatically returns zero values (e.g., "", 0, nil etc.).
func Todo() {
	m := management{}
	targetSkip, sourceSkip := m.skipNums()
	target := m.funcInfo(targetSkip)
	source := m.funcInfo(sourceSkip)
	msg := todoFormat(target, source)
	if event != nil {
		event()
		return
	}
	fmt.Println(msg)
}

func todoFormat(target, source string) string {
	strs := make([]string, 5)
	strs = append(strs, "attribute.Todo: todo in ", target, "; ")
	strs = append(strs, "source: ", source)
	res := strings.Join(strs, "")
	return res
}
