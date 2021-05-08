package _1_简化路径

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	nums := strings.Split(path, "/")
	var stack []string

	for _, s := range nums {
		if s == ".." {
			if len(stack) != 0 {
				stack = stack[:len(stack)-1]
			}
		} else if s != "." && s != "" {
			stack = append(stack, s)
		}
	}

	fmt.Println(stack)

	var result string
	if len(stack) != 0 {
		for _, s := range stack {
			result += "/" + s
		}
		return result
	}

	return "/"

}
