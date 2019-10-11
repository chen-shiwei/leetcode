package _226

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInvertTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		//{
		//	name: "[4,2,7,1,3,6,9]=>[4,7,2,9,6,3,1]",
		//	args: args{root: &TreeNode{
		//		Val: 4,
		//		Left: &TreeNode{
		//			Val: 2,
		//			Left: &TreeNode{
		//				Val:   1,
		//				Left:  nil,
		//				Right: nil,
		//			},
		//			Right: &TreeNode{
		//				Val:   3,
		//				Left:  nil,
		//				Right: nil,
		//			},
		//		},
		//		Right: &TreeNode{
		//			Val: 7,
		//			Left: &TreeNode{
		//				Val:   6,
		//				Left:  nil,
		//				Right: nil,
		//			},
		//			Right: &TreeNode{
		//				Val:   9,
		//				Left:  nil,
		//				Right: nil,
		//			},
		//		},
		//	}},
		//	want: &TreeNode{
		//		Val: 4,
		//		Left: &TreeNode{
		//			Val: 7,
		//			Left: &TreeNode{
		//				Val:   9,
		//				Left:  nil,
		//				Right: nil,
		//			},
		//			Right: &TreeNode{
		//				Val:   6,
		//				Left:  nil,
		//				Right: nil,
		//			},
		//		},
		//		Right: &TreeNode{
		//			Val: 2,
		//			Left: &TreeNode{
		//				Val:   3,
		//				Left:  nil,
		//				Right: nil,
		//			},
		//			Right: &TreeNode{
		//				Val:   1,
		//				Left:  nil,
		//				Right: nil,
		//			},
		//		}},
		//},
		{
			name: "[1,2]=>[1,null,2]",
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			}},
			want: &TreeNode{
				Val:  1,
				Left: nil,
				Right: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InvertTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				fmt.Println("got", got.Left, got.Right)
				fmt.Println("want", tt.want.Left, tt.want.Right)
				t.Errorf("InvertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
