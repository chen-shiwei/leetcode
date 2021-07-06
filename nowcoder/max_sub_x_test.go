package nowcoder

import "testing"

func Test_getMaxSubX(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: ``, args: args{nums: []int{-2, -1, 0, 1, 2, 0, -1, -2, 5, 4, 6, 7}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaxSubX(tt.args.nums); got != tt.want {
				t.Errorf("getMaxSubX() = %v, want %v", got, tt.want)
			}
		})
	}
}
