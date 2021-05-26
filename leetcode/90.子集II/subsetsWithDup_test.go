package _0_子集II

import (
	"reflect"
	"testing"
)

func Test_subsetsWithDup(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: `输入：nums = [1,2,2]
输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]`, args: args{nums: []int{1, 2, 2}}, want: [][]int{
			{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsetsWithDup(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subsetsWithDup() = %v, want %v", got, tt.want)
			}
		})
	}
}
