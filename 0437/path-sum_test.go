package _437

import "testing"

func TestPathSum(t *testing.T) {
	type args struct {
		root *TreeNode
		sum  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PathSum(tt.args.root, tt.args.sum); got != tt.want {
				t.Errorf("PathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
