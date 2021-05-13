package _151_翻转字符串里的单词

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: `输入："the sky is blue"
			输出："blue is sky the"`, args: args{s: "the sky is blue"}, want: "blue is sky the"},
		{name: `输入："  hello world!  "
输出："world! hello"`, args: args{s: "  hello world!  "}, want: "world! hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func Test(t *testing.T) {

	a := []byte("abcd")

	fmt.Println(a[0])
	fmt.Println(string(a[:2]))
	fmt.Println(string(a[1:3]))
}

func Test_removeSpaces1(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "", args: args{b: []byte(" a b c ")}, want: []byte("a b c")},
		{name: "", args: args{b: []byte(" a b c")}, want: []byte("a b c")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeSpaces(tt.args.b)

			fmt.Println(tt.args.b)
			fmt.Println(got)
			fmt.Println(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}
