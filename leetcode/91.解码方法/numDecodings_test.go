package _1_解码方法

import (
	"fmt"
	"testing"
)

func Test_numDecodings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: `输入：s = "12"
输出：2`, args: args{s: "12"}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDecodings(tt.args.s); got != tt.want {
				t.Errorf("numDecodings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test(t *testing.T) {
	var s = "A"
	fmt.Println(s[0] - '0')
}
