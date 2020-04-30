package _095

import (
	"testing"
)

func Test_findInMountainArray(t *testing.T) {
	type args struct {
		target      int
		mountainArr *MountainArray
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `示例 1：
				输入：array = [1,2,3,4,5,3,1], target = 3
				输出：2
				解释：3 在数组中出现了两次，下标分别为 2 和 5，我们返回最小的下标 2。
				
				来源：力扣（LeetCode）
				链接：https://leetcode-cn.com/problems/find-in-mountain-array
				著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。`,
			args: args{
				target:      3,
				mountainArr: NewMountainArray([]int{1, 2, 3, 4, 5, 3, 1}),
			},
			want: 2},
		{
			name: `示例 2：
					输入：array = [0,1,2,4,2,1], target = 3
					输出：-1
					解释：3 在数组中没有出现，返回 -1。`,
			args: args{
				target:      3,
				mountainArr: NewMountainArray([]int{0, 1, 2, 4, 2, 1}),
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findInMountainArray(tt.args.target, tt.args.mountainArr); got != tt.want {
				t.Errorf("findInMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
