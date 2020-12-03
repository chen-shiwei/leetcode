package _04_计数质数

import "testing"

func Test_isPrimes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: `5`, args: args{n: 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPrimes(tt.args.n); got != tt.want {
				t.Errorf("isPrimes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countPrimes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: `输入：n = 10
输出：4
解释：小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。`, args: args{n: 10}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPrimes(tt.args.n); got != tt.want {
				t.Errorf("countPrimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
