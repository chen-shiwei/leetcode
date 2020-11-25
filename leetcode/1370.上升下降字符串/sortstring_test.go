package _370_上升下降字符串

import (
	"fmt"
	"testing"
)

func Test_getMinChar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: `abc,c`, args: args{s: "abc"}, want: "a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, got := getMinChar(tt.args.s); string(got) != tt.want {
				t.Errorf("getMinChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMaxChar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: `abc,c`, args: args{s: "abc"}, want: "c"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, got := getMaxChar(tt.args.s); string(got) != tt.want {
				t.Errorf("getMaxChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeIdxStr(t *testing.T) {
	type args struct {
		s   string
		idx int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: `abc`, args: args{
			s:   "abc",
			idx: 0,
		}, want: "bc"},
		{name: `abc`, args: args{
			s:   "abc",
			idx: 2,
		}, want: "ab"},
		{name: `a`, args: args{
			s:   "a",
			idx: 0,
		}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeIdxStr(tt.args.s, tt.args.idx)
			if got != tt.want {
				t.Errorf("removeIdxStr() = %v, want %v", got, tt.want)
			}
			fmt.Println(tt.args.s, tt.args.idx, got)

		})
	}
}

func Test_sortString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: `输入：s = "aaaabbbbcccc"
输出："abccbaabccba"`, args: args{s: "aaaabbbbcccc"}, want: "abccbaabccba"},
		{name: `输入：s = "rat"
输出："art"
解释：单词 "rat" 在上述算法重排序以后变成 "art"`, args: args{s: "rat"}, want: "art"},
		{name: `输入：s = "leetcode"
输出："cdelotee"`, args: args{s: "leetcode"}, want: "cdelotee"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortString(tt.args.s); got != tt.want {
				t.Errorf("sortString() = %v, want %v", got, tt.want)
			}
		})
	}
}
