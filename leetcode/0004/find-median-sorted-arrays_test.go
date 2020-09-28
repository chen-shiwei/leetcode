package _004

import (
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "13 2 => 2", args: args{
			nums1: []int{1, 3},
			nums2: []int{2},
		}, want: 2},
		{name: "12 34 => 2.5", args: args{
			nums1: []int{1, 2},
			nums2: []int{3, 4},
		}, want: 2.5},
		{name: "1 1 => 1.0", args: args{
			nums1: []int{1},
			nums2: []int{1},
		}, want: 1.0}, {name: "12 12 => 1.5", args: args{
			nums1: []int{1, 2},
			nums2: []int{1, 2},
		}, want: 1.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("FindMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
