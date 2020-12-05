package _5_颜色分类

import "testing"

func Test_sortColors(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: `输入：nums = [2,0,2,1,1,0]
输出：[0,0,1,1,2,2]`, args: args{nums: []int{2, 0, 2, 1, 1, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.args.nums)

			t.Logf("%+v", tt.args.nums)
		})
	}
}
