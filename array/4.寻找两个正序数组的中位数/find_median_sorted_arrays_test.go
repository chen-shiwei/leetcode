package __寻找两个正序数组的中位数

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: `
示例 1：
输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2`, args: args{
			nums1: []int{1, 3},
			nums2: []int{2},
		}, want: 2},
		{name: `示例 2：
输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5`, args: args{
			nums1: []int{1, 2},
			nums2: []int{3, 4},
		}, want: 2.5},
		{name: `示例 3：
输入：nums1 = [0,0], nums2 = [0,0]
输出：0.00000`, args: args{
			nums1: []int{0, 0},
			nums2: []int{0, 0},
		}, want: 0.0},
		{name: `示例 5：
输入：nums1 = [2], nums2 = []
输出：2.00000`, args: args{
			nums1: []int{2},
			nums2: nil,
		}, want: 2.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays1(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays1() = %v, want %v", got, tt.want)
			}
		})
	}
}
