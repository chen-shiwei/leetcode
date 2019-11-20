package _394_todo

import "testing"

func TestDecodeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: `s = "3[a]2[bc]", 返回 "aaabcbc".`, args: args{s: "3[a]2[bc]"}, want: "aaabcbc"},
		{name: `s = "3[a2[c]]", 返回 "accaccacc".`, args: args{s: "3[a2[c]]"}, want: "accaccacc"},
		{name: `s = "2[abc]3[cd]ef", 返回 "abcabccdcdcdef".`, args: args{s: "2[abc]3[cd]ef"}, want: "abcabccdcdcdef"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecodeString(tt.args.s); got != tt.want {
				t.Errorf("DecodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
