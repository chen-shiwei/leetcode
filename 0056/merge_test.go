package _056

import (
	"reflect"
	"testing"
)

//
func TestMerge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: `输入: [[1,3],[2,6],[8,10],[15,18]]
				输出: [[1,6],[8,10],[15,18]]
				解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].`,
			args: args{intervals: [][]int{{2, 6}, {1, 3}, {8, 10}, {15, 18}}},
			want: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name: `输入: [[1,4],[4,5]]
				输出: [[1,5]]
				解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。`,
			args: args{intervals: [][]int{{1, 4}, {4, 5}}},
			want: [][]int{{1, 5}},
		},
		{
			name: `输入: [[1,4],[0,4]]
				输出: [[1,4]]
				解释: 区间 [1,4] 和 [0,4] 可被视为重叠区间。`,
			args: args{intervals: [][]int{{1, 4}, {0, 4}}},
			want: [][]int{{0, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
