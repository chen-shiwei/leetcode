package _59_重复的子字符串

import "testing"

func Test_repeatedSubstringPattern(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: `输入: "abab"`, args: args{s: "abab"}, want: true},
		{name: `asdfasdfasdf`, args: args{s: "asdfasdfasdf"}, want: true},
		{name: `aba`, args: args{s: "aba"}, want: false},
		{name: `abcabcabcabc`, args: args{s: "abcabcabcabc"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedSubstringPattern(tt.args.s); got != tt.want {
				t.Errorf("repeatedSubstringPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
