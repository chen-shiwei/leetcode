package _47_前K个高频元素

import (
	"reflect"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: `输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]`, args: args{
			nums: []int{1, 1, 1, 2, 2, 3},
			k:    2,
		}, want: []int{1, 2}},
		{name: `[5,2,5,3,5,3,1,1,3]
2`, args: args{
			nums: []int{5, 2, 5, 3, 5, 3, 1, 1, 3},
			k:    2,
		}, want: []int{3, 5}},
		{name: `[1,1,1,2,2,3333]
2`, args: args{
			nums: []int{1, 1, 1, 2, 2, 3333},
			k:    2,
		}, want: []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequent(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
