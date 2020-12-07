package _52_乘积最大子数组

import "testing"

func Test_maxProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: `输入: [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。`, args: args{nums: []int{2, 3, -2, 4}}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.args.nums); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
