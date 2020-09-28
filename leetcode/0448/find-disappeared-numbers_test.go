package _448

import (
	"reflect"
	"testing"
)

func TestFindDisappearedNumbers(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: `输入:
				[4,3,2,7,8,2,3,1]
				
				输出:
				[5,6]`, args: args{nums: []int{4, 3, 2, 7, 8, 2, 3, 1}}, want: []int{5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindDisappearedNumbers(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindDisappearedNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
