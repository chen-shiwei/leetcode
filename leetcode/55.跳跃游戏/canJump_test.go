package _5_跳跃游戏

import "testing"

func Test_canJump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: `输入：nums = [2,3,1,1,4],输出：true`, args: args{nums: []int{2, 3, 1, 1, 4}}, want: true},
		{name: `输入：nums = [3,2,1,0,4],输出：false`, args: args{nums: []int{3, 2, 1, 0, 4}}, want: false},
		{name: `输入：nums = [2,0,0],输出：trur`, args: args{nums: []int{2, 0, 0}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump(tt.args.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}
