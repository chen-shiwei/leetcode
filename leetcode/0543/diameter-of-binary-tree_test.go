package _543

import "testing"

func TestDiameterOfBinaryTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "[1,2,3,4,5]", args: args{root: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
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
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DiameterOfBinaryTree(tt.args.root); got != tt.want {
				t.Errorf("DiameterOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test1(t *testing.T) {
	println("1")
}
