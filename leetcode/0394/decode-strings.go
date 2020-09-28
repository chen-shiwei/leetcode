package _394

import (
	"bytes"
	"github.com/chen-shiwei/leetcode/datastruct"
	"strconv"
	"unicode"
)

// DecodeString
// https://leetcode-cn.com/problems/decode-string
func DecodeString(s string) string {

	l := len(s)

	wordStack := datastruct.NewStack()
	multiStack := datastruct.NewStack()
	var n int
	var result, tmp string
	for i := 0; i < l; i++ {
		word, _ := strconv.Atoi(string(s[i]))
		if unicode.IsDigit(rune(s[i])) {
			if multiStack.Empty() {
				result += tmp
				tmp = ""
			}
			n = n*10 + word
		} else if s[i] == '[' {
			multiStack.Push(n)
			wordStack.Push(tmp)
			n = 0
			tmp = ""
		} else if s[i] == ']' {
			n := multiStack.Pop().(int)
			if !multiStack.Empty() {
				tmp = wordStack.Pop().(string) + nString(n, tmp)
			} else {
				result += nString(n, tmp)
				tmp = ""
			}
		} else {
			tmp += string(s[i])
		}
	}
	return result + tmp
}

func nString(n int, word string) string {
	var s bytes.Buffer
	for i := 0; i < n; i++ {
		s.WriteString(word)
	}
	return s.String()
}
