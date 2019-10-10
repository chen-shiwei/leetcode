package _136

import "testing"

func TestSingleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{args: args{nums: []int{1, 2, 3, 3, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SingleNumber(tt.args.nums); got != tt.want {
				t.Errorf("SingleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
