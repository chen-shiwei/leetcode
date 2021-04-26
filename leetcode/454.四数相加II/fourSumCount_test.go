package _54_四数相加II

import "testing"

func Test_fourSumCount(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		nums3 []int
		nums4 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: `输入:
A = [ 1, 2]
B = [-2,-1]
C = [-1, 2]
D = [ 0, 2]

输出:
2

解释:
两个元组如下:
1. (0, 0, 0, 1) -> A[0] + B[0] + C[0] + D[1] = 1 + (-2) + (-1) + 2 = 0
2. (1, 1, 0, 0) -> A[1] + B[1] + C[0] + D[0] = 2 + (-1) + (-1) + 0 = 0
`, args: args{
			nums1: []int{1, 2},
			nums2: []int{-2, -1},
			nums3: []int{-1, 2},
			nums4: []int{0, 2},
		}, want: 2},
		{name: `[-1,-1]
[-1,1]
[-1,1]
[1,-1]`, args: args{
			nums1: []int{-1, -1},
			nums2: []int{-1, 1},
			nums3: []int{-1, 1},
			nums4: []int{1, -1},
		}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fourSumCount(tt.args.nums1, tt.args.nums2, tt.args.nums3, tt.args.nums4); got != tt.want {
				t.Errorf("fourSumCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
