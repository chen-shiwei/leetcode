package _101

import "testing"

func TestIsSymmetric(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "[1,0]", args: args{root: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		}}, want: false},
		{name: "[1,2,3]", args: args{root: &TreeNode{
			Val: 1,
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
		}}, want: false},
		{name: "[2,3,3,4,5,null,4]", args: args{root: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &TreeNode{
				Val:  3,
				Left: nil,
				Right: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
			},
		}}, want: false}, {name: "[2,3,3,4,5,5]", args: args{root: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
		}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSymmetric(tt.args.root); got != tt.want {
				t.Errorf("IsSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
