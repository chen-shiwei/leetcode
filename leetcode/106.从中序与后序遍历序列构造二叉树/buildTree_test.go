package _06_从中序与后序遍历序列构造二叉树

import (
	"log"
	"reflect"
	"testing"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		inorder   []int
		postorder []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{name: `中序遍历 inorder = [9,3,15,20,7]
后序遍历 postorder = [9,15,7,20,3]`, args: struct {
			inorder   []int
			postorder []int
		}{inorder: []int{9, 3, 15, 20, 7}, postorder: []int{9, 15, 7, 20, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildTree(tt.args.inorder, tt.args.postorder)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}

			InEach(got)
		})
	}
}

func InEach(root *TreeNode) {
	if root == nil {
		return
	}
	InEach(root.Left)
	log.Println(root.Val)
	InEach(root.Right)

	return
}
