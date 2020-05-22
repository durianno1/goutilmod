package goutilmod

import (
	"fmt"
)

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func Hi(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}
