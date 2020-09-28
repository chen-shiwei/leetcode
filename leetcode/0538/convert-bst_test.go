package _538

import (
	"reflect"
	"testing"
)

func TestConvertBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{name: "[5,2,13]=>[18,20,13]", args: args{root: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   13,
				Left:  nil,
				Right: nil,
			},
		}}, want: &TreeNode{
			Val: 18,
			Left: &TreeNode{
				Val:   20,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   13,
				Left:  nil,
				Right: nil,
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertBST(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
