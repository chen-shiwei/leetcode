package _542

import (
	"reflect"
	"testing"
)

func Test_updateMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: `
输入: 
0 0 0
0 1 0
0 0 0
输出:
0 0 0
0 1 0
0 0 0`, args: args{matrix: [][]int{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
		}}, want: [][]int{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
		}},
		{name: `
输入:
0 0 0
0 1 0
1 1 1
输出:
0 0 0
0 1 0
1 2 1`, args: args{matrix: [][]int{
			{0, 0, 0},
			{0, 1, 0},
			{1, 1, 1},
		}}, want: [][]int{
			{0, 0, 0},
			{0, 1, 0},
			{1, 2, 1},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := updateMatrix(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
