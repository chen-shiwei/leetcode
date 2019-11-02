package _169

import "testing"

func TestMajorityElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: `"输入: [3,2,3] 输出: 3`, args: args{[]int{3, 2, 3}}, want: 3},
		{name: `"输入: [1] 输出: 1`, args: args{[]int{1}}, want: 1},
		{name: `"输入: [2,2,1,1,1,2,2] 输出: 2`, args: args{[]int{2, 2, 1, 1, 1, 2, 2}}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MajorityElement(tt.args.nums); got != tt.want {
				t.Errorf("MajorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
