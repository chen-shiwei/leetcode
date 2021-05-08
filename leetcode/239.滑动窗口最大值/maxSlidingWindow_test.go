package _39_滑动窗口最大值

import (
	"reflect"
	"testing"
)

func Test_maxSlidingWindow(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: `输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]`, args: args{
			nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:    3,
		}, want: []int{3, 3, 5, 5, 6, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSlidingWindow1(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSlidingWindow1() = %v, want %v", got, tt.want)
			}
		})
	}
}
