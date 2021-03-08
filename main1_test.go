package main

import "testing"

func TestFind(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: ``, args: args{
			nums:   []int{1, 2, 2, 3, 3, 5, 6, 9, 9},
			target: 4,
		}, want: 5},

		{name: ``, args: args{
			nums:   []int{1, 2, 2, 3, 3, 5, 6, 9, 9},
			target: 1,
		}, want: 0},
		{name: ``, args: args{
			nums:   []int{1, 2, 2, 3, 3, 5, 6, 9, 9},
			target: 10,
		}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Find(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
