package _005_K次取反后最大化的数组和

import "testing"

func Test_largestSumAfterKNegations(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: `[3,-1,0,2]
3`, args: args{
			nums: []int{3, -1, 0, 2},
			k:    3,
		}, want: 6},
		{name: `[2,-3,-1,5,-4]
2`, args: args{
			nums: []int{2, -3, -1, 5, -4},
			k:    2,
		}, want: 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestSumAfterKNegations(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("largestSumAfterKNegations() = %v, want %v", got, tt.want)
			}
		})
	}
}
