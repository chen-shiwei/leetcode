package _287

import "testing"

func TestFindDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: `输入: [1,3,4,2,2]
				输出: 2`, args: args{nums: []int{1, 3, 4, 2, 2}}, want: 2},
		{name: `输入: [3,1,3,4,2]
				输出: 3`, args: args{nums: []int{3, 1, 3, 4, 2}}, want: 3},
		{name: `输入: [8,7,1,10,17,15,18,11,16,9,19,12,5,14,3,4,2,13,18,18]
				输出: 18`, args: args{nums: []int{8, 7, 1, 10, 17, 15, 18, 11, 16, 9, 19, 12, 5, 14, 3, 4, 2, 13, 18, 18}}, want: 18},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("FindDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
