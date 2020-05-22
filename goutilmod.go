package goutilmod

import (
	"fmt"
)

func Hi(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
