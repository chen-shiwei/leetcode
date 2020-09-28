package _248

import "testing"

func Test_numOfSubArrays(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: ` 
				示例 1：
				输入：nums = [{1,1,2,1,1}], k = 3
				输出：2
				解释：包含 3 个奇数的子数组是 [1,1,2,1] 和 [1,2,1,1] 。
				`, args: args{
			nums: []int{1, 1, 2, 1, 1},
			k:    3,
		}, want: 2},
		{name: `示例 2：
				输入：nums = [2,4,6], k = 1
				输出：0
				解释：数列中不包含任何奇数，所以不存在优美子数组。`, args: args{
			nums: []int{2, 4, 6},
			k:    1,
		}, want: 0},
		{name: `示例 3：
				输入：nums = [2,2,2,1,2,2,1,2,2,2], k = 2
				输出：16`, args: args{
			nums: []int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2},
			k:    2,
		}, want: 16},
	}
	for _, tt := range tests {
		t.Run(tt.name+" numOfSubArrays", func(t *testing.T) {
			if got := numOfSubArrays(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("numOfSubArrays() = %v, want %v", got, tt.want)
			}

		})

	}
}
