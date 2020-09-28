package sort

import (
	"sort"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	type args struct {
		nums sort.Interface
		desc bool
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "1", args: args{nums: sort.IntSlice{22, 34, 3, 40, 18, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//BubbleSort(tt.args.nums, tt.args.desc)
			bubbleSort(tt.args.nums, tt.args.desc)
			t.Logf("nums:%v", tt.args.nums)
		})
	}
}
