package _52_用最少数量的箭引爆气球

import "testing"

func Test_findMinArrowShots(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: `输入：points = [[10,16],[2,8],[1,6],[7,12]]
输出：2`, args: args{points: [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinArrowShots(tt.args.points); got != tt.want {
				t.Errorf("findMinArrowShots() = %v, want %v", got, tt.want)
			}
		})
	}
}
