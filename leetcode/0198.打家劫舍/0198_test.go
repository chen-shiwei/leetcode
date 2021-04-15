package _198_打家劫舍

import "testing"

func Test_rob(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: `输入: [1,2,3,1]
				输出: 4
				解释: 偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
				     偷窃到的最高金额 = 1 + 3 = 4 。`, args: args{nums: []int{1, 2, 3, 1}}, want: 4},
		{name: `输入: [2,7,9,3,1]
				输出: 12
				解释: 偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
				     偷窃到的最高金额 = 2 + 9 + 1 = 12 。`, args: args{nums: []int{2, 7, 9, 3, 1}}, want: 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
			if got := robDP(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
