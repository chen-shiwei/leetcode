package 判断小数组是否是大数组的连续子集

import "testing"

func Test_isArraySub(t *testing.T) {
	type args struct {
		bigNums   []int
		smallNums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "", args: args{
			bigNums:   []int{1, 3, 5, 6, 7, 8, 9},
			smallNums: []int{7, 8, 9},
		}, want: 4},
		{name: "", args: args{
			bigNums:   []int{1, 3, 7, 8, 5, 6, 7, 8, 9},
			smallNums: []int{7, 8, 9},
		}, want: 6},
		{name: "", args: args{
			bigNums:   []int{1, 3, 7, 8, 5, 6, 7, 8},
			smallNums: []int{7, 8, 9},
		}, want: -1},
		{name: "", args: args{
			bigNums:   []int{7, 8, 9},
			smallNums: []int{7, 8, 9},
		}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isArraySub(tt.args.bigNums, tt.args.smallNums); got != tt.want {
				t.Errorf("isArraySub() = %v, want %v", got, tt.want)
			}
		})
	}
}
