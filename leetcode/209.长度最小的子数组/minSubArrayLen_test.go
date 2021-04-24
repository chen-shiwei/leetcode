package _09_长度最小的子数组

import "testing"

func Test_minSubArrayLenViolence(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: `输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。`, args: args{
			target: 7,
			nums:   []int{2, 3, 1, 2, 4, 3},
		}, want: 2},
		{name: `输入：target = 4, nums = [1,4,4]
输出：1`, args: args{
			target: 4,
			nums:   []int{1, 4, 4},
		}, want: 1},
		{name: `输入：target = 11, nums = [1,1,1,1,1,1,1,1]
输出：0`, args: args{
			target: 11,
			nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
		}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLenViolence(tt.args.target, tt.args.nums); got != tt.want {
				t.Errorf("minSubArrayLenViolence() = %v, want %v", got, tt.want)
			}
			if got := minSubArrayLenSlidingWindow(tt.args.target, tt.args.nums); got != tt.want {
				t.Errorf("minSubArrayLenSlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
