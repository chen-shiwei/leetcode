package _1_下一个排列

import "testing"

func Test_nextPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: `123465`, args: args{nums: []int{1, 2, 3, 4, 6, 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextPermutation(tt.args.nums)
			t.Log(tt.args.nums)
		})
	}
}
