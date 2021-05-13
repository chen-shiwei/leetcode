package __无重复字符的最长子串

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: `输入: s = "abcabcbb"输出: 3解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。`, args: args{s: "abcabcbb"}, want: 3},
		{name: `输入: s = "dvdf"输出: 3`, args: args{s: "dvdf"}, want: 3},
		{name: `输入: s = "au"输出: 2`, args: args{s: "au"}, want: 2},
		{name: `输入: s = ""abcabcbb""输出: 4`, args: args{s: "abcabcbb"}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
