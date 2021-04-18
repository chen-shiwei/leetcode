package sort

import (
	"sort"
	"testing"
)

func Test_selectSort(t *testing.T) {
	type args struct {
		nums sort.Interface
		desc bool
	}
	tests := []struct {
		name string
		args args
	}{
		{name: `asc`, args: args{
			nums: sort.IntSlice{22, 34, 3, 40, 18, 4},
			desc: false,
		}},
		{name: `desc`, args: args{
			nums: sort.IntSlice{22, 34, 3, 40, 18, 4},
			desc: true,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log("before", tt.args.nums)
			selectSort(tt.args.nums, tt.args.desc)
			t.Log("after", tt.args.nums)
		})
	}
}
