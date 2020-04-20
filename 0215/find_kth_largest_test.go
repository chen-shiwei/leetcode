package _215

import "testing"

func TestFindKthLargest(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: `输入: [3,2,1,5,6,4] 和 k = 2
输出: 5`, args: args{
			nums: []int{3, 2, 1, 5, 6, 4},
			k:    2,
		}, want: 5},
		{
			name: `输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
输出: 4`, args: args{
				nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
				k:    4,
			}, want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindKthLargest(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("FindKthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}
