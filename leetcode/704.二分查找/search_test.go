package _04_二分查找

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: `输入: nums = [-1,0,3,5,9,12], target = 9
输出: 4
解释: 9 出现在 nums 中并且下标为 4`, args: args{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 9,
		}, want: 4},
		{name: `输入: nums = [-1,0,3,5,9,12], target = 2
输出: -1
解释: 2 不存在 nums 中因此返回 -1`, args: args{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 2,
		}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
			if got := search1(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search1() = %v, want %v", got, tt.want)
			}
		})
	}
}
