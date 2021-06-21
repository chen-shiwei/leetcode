package _746_使用最小花费爬楼梯

import "testing"

func Test_minCostClimbingStairs(t *testing.T) {
	type args struct {
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: `输入：cost = [10, 15, 20] 输出：15`, args: args{cost: []int{10, 15, 20}}, want: 15},
		{name: `输入：cost = [1,100,1,1,1,100,1,1,100,1] 输出：15`, args: args{cost: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostClimbingStairs(tt.args.cost); got != tt.want {
				t.Errorf("minCostClimbingStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
