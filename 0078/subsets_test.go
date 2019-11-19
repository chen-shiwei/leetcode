package _078

import (
	"reflect"
	"testing"
)

func TestSubsets(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: `输入: nums = [1,2,3]
				输出:
				[
				  [3],
				  [1],
				  [2],
				  [1,2,3],
				  [1,3],
				  [2,3],
				  [1,2],
				  []
				]`,
			args: args{nums: []int{1, 2, 3, 4}},
			want: [][]int{{3}, {1}, {2}, {1, 2, 3}, {1, 3}, {2, 3}, {1, 2}, {}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Subsets(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
