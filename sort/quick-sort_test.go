package sort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "55, 94, 87, 1, 4, 32, 11, 77, 39, 42, 64, 53, 70, 12, 9,", args: args{nums: []int{55, 94, 87, 1, 4, 32, 11, 77, 39, 42, 64, 53, 70, 12, 9}}, want: []int{1, 4, 9, 11, 12, 32, 39, 42, 53, 55, 64, 70, 77, 87, 94}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var got = make([]int, len(tt.args.nums))

			QuickSort(got)

			t.Log(tt.args.nums, got, tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
