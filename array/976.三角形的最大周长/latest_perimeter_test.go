package _76_三角形的最大周长

import "testing"

func Test_largestPerimeter(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: `[1,2,2,4,18,8]`, args: args{
			A: []int{1, 2, 2, 4, 18, 8},
		}, want: 5},
		{name: `3,4,15,2,9,4`, args: args{A: []int{3, 4, 15, 2, 9, 4}}, want: 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestPerimeter(tt.args.A); got != tt.want {
				t.Errorf("largestPerimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}
