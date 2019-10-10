package sort

import "testing"

func TestBinarySort(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name  string
		args  args
		wantI int
	}{
		// TODO: Add test cases.
		//Input : [1,2,3,4,5]
		//key : 3
		//return the index : 2
		{args: args{
			nums: []int{1, 2, 3, 4, 5},
			val:  3,
		}, wantI: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotI := BinarySort(tt.args.nums, tt.args.val); gotI != tt.wantI {
				t.Errorf("BinarySort() = %v, want %v", gotI, tt.wantI)
			}
		})
	}
}
