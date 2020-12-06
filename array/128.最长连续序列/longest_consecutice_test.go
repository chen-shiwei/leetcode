package _28_最长连续序列

import "testing"

func Test_longestConsecutive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: `输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。`, args: args{nums: []int{100, 4, 200, 1, 3, 2}}, want: 4},
		{name: `输入：nums = [0,3,7,2,5,8,4,6,0,1]
输出：9`, args: args{nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}}, want: 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestConsecutive(tt.args.nums); got != tt.want {
				t.Errorf("longestConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}
