package _5_三数之和

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "[-1, 0, 1, 2, -1, -4]", args: args{nums: []int{-1, 0, 1, 2, -1, -4}}, want: [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{name: "[0,0,0]", args: args{nums: []int{0, 0, 0}}, want: [][]int{{0, 0, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ThreeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ThreeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
