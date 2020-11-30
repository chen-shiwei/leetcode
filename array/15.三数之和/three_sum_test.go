package _5_三数之和

import (
	"reflect"
	"testing"
)

func Test_threeSum1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: `给定数组 nums = [-1, 0, 1, 2, -1, -4]，
满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]`, args: args{nums: []int{
			-1, 0, 1, 2, -1, -4}}, want: [][]int{
			{-1, 0, 1},
			{-1, -1, 2},
		}},
		{name: `[0,0,0]`, args: args{nums: []int{0, 0, 0}}, want: [][]int{{0, 0, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum1(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum1() = %v, want %v", got, tt.want)
			}
		})
	}
}
