package _4_柱状图中最大的矩形

import "testing"

func Test_largestRectangleArea(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: `输入: [2,1,5,6,2,3]
输出: 10`, args: args{heights: []int{2, 1, 5, 6, 2, 3}}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestRectangleArea(tt.args.heights); got != tt.want {
				t.Errorf("largestRectangleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
