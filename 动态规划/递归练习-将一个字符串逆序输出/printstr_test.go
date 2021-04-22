package 递归练习_将一个字符串逆序输出

import "testing"

func Test_reversePrintStr(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: `abcdef`, args: args{s: "abcdef"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reversePrintStr(tt.args.s)
		})
	}
}
