package _8_旋转图像

import "testing"

func Test_rotate(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: `给定 matrix = 
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],

原地旋转输入矩阵，使其变为:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]`, args: args{matrix: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}}, want: [][]int{
			{7, 4, 1},
			{8, 5, 2},
			{9, 6, 3},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("befor:%+v", tt.args.matrix)
			rotate(tt.args.matrix)
			t.Logf("after:%+v", tt.args.matrix)
		})
	}
}
