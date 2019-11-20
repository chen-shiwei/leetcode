package _394_todo

import (
	"github.com/chen-shiwei/leetcode/datastruct"
	"strconv"
)

var (
	leftFlag  = "]"
	rightFlag = "]"
)

func DecodeString(s string) string {

	l := len(s)

	stack := datastruct.NewStack()
	//r := []rune(s)
	//l := len(r)
	//
	////fmt.Println(string(s[0]))
	//var n int
	for i := 0; i < l; i++ {
		word, _ := strconv.Atoi(string(s[i]))
		if word > 0 {
			n = word
		}
		if s[i] == '[' {

			break
		}
		if s[i] == ']' {
		}
		stack.Push(s[i])
	}
	//return ""
}

func join(w rune) string {

	n, _ := strconv.Atoi(string(w))
	if n > 0 {
		for i := 0; i < l; i++ {
			s += join(l, i+1, words)
		}
	}
	if s[i] == '[' {

	}
	if s[i] == ']' {

	}

	return w
}
