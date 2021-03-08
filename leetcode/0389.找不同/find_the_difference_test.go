package _389_找不同

import "testing"

func Test_findTheDifference(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{name: `示例 1：

输入：s = "abcd", t = "abcde"
输出："e"
解释：'e' 是那个被添加的字母。`, args: args{
			s: "abcd",
			t: "abcde",
		}, want: 'e'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheDifference(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("findTheDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
