package sort

import (
	"reflect"
	"sort"
	"testing"
)

func TestInsertSort1(t *testing.T) {
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
			if got := InsertSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_insertSort(t *testing.T) {
	type args struct {
		nums sort.Interface
		desc bool
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "55, 94, 87, 1, 4, 32, 11, 77, 39, 42, 64, 53, 70, 12, 9,", args: args{nums: sort.IntSlice{55, 94, 87, 1, 4, 32, 11, 77, 39, 42, 64, 53, 70, 12, 9}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log("before", tt.args.nums)
			insertSort(tt.args.nums, tt.args.desc)
			t.Log("after", tt.args.nums)
		})
	}
}
