package offer_58

import "testing"

func Test_reverseLeftWords(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: `输入: s = "abcdefg", k = 2
输出: "cdefgab"`, args: args{
			s: "abcdefg",
			n: 2,
		}, want: "cdefgab"},
		{name: `输入: s = "lrloseumgh", k = 6
输出: "umghlrlose"`, args: args{
			s: "lrloseumgh",
			n: 6,
		}, want: "umghlrlose"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseLeftWords(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("reverseLeftWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
