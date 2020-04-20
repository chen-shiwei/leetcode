package _114

import (
	"reflect"
	"testing"
)

func Test_flatten(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{name: `1
			   / \
			  2   5
			 / \   \
			3   4   6
			展开为:
			1
			 \
			  2
			   \
				3
				 \
				  4
				   \
					5
					 \
					  6`, args: args{root: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &TreeNode{
				Val:  5,
				Left: nil,
				Right: &TreeNode{
					Val:   6,
					Left:  nil,
					Right: nil,
				},
			},
		}}, want: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val:  2,
				Left: nil,
				Right: &TreeNode{
					Val:  3,
					Left: nil,
					Right: &TreeNode{
						Val:  4,
						Left: nil,
						Right: &TreeNode{
							Val:  5,
							Left: nil,
							Right: &TreeNode{
								Val:   6,
								Left:  nil,
								Right: nil,
							},
						},
					},
				},
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flatten(tt.args.root)
			if !reflect.DeepEqual(tt.args.root, tt.want) {
				t.Fail()
			}
		})
	}
}
