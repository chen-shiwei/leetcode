package _44_比较含退格的字符串

import "testing"

func Test_backspaceCompare(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: `输入：S = "ab#c", T = "ad#c"
输出：true
解释：S 和 T 都会变成 “ac”。`, args: args{
			s: "ab#c",
			t: "ad#c",
		}, want: true},
		{name: `输入：S = "ab##", T = "c#d#"
输出：true
解释：S 和 T 都会变成 “”。`, args: args{
			s: "ab##",
			t: "c#d#",
		}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backspaceCompare(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("backspaceCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}
