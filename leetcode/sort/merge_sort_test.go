package sort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "9,8,7,6,5,4,3,2,1", args: args{nums: []int{9, 8, 7, 6, 5, 4, 3, 2, 1}},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeSort(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Fatalf("\ngot:%v \nwant:%v", tt.args.nums, tt.want)
			}

		})
	}
}
