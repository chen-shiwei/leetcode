package _098

import (
	"fmt"
	"math"
	"testing"
)

func TestIsValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "[2,1,3]", args: args{root: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		}}, want: true},
		{name: "[5,1,4,null,null,3,6]", args: args{root: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   6,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   6,
					Left:  nil,
					Right: nil,
				},
			},
		}}, want: false},
		{name: "[10,5,15,null,null,6,20]", args: args{root: &TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 15,
				Left: &TreeNode{
					Val:   6,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   20,
					Left:  nil,
					Right: nil,
				},
			},
		}}, want: false},
		{name: "[0]", args: args{root: &TreeNode{
			Val:   0,
			Left:  nil,
			Right: nil,
		}}, want: true},
		{name: "[]", args: args{root: nil}, want: true},
		{name: "[-2147483648]", args: args{root: &TreeNode{
			Val:   -2147483648,
			Left:  nil,
			Right: nil,
		}}, want: true},
		{name: "[1,1]", args: args{root: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidBST(tt.args.root); got != tt.want {
				t.Errorf("IsValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test(t *testing.T) {
	fmt.Println(-2147483648, math.MinInt32-1)

}
