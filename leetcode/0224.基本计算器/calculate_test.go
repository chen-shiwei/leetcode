package _224_基本计算器

import "testing"

func Test_calculate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: `"1 + 1"`, args: args{s: "1 + 1"}, want: 2},
		{name: `"(1+(4+5+2)-3)+(6+8)"`, args: args{s: "(1+(4+5+2)-3)+(6+8)"}, want: 23},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.args.s); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
