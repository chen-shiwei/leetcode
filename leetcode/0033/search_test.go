package _033

import "testing"

func TestSearch(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: `输入: nums = [4,5,6,7,0,1,2], target = 0
				输出: 4`, args: args{
			nums:   []int{4, 5, 6, 7, 0, 1},
			target: 0,
		}, want: 4},
		{name: `输入: nums = [4,5,6,7,0,1,2], target = 3
				输出: -1`, args: args{
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
