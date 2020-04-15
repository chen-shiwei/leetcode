package _102

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "[3,9,20,null,null,15,7]", args: args{root: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 20,
				Left: &TreeNode{
					Val:   15,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
			},
		}}, want: [][]int{{3}, {9, 20}, {15, 7}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LevelOrder1(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LevelOrder1() = %v, want %v", got, tt.want)
			}
			if got := LevelOrder2(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LevelOrder2() = %v, want %v", got, tt.want)
			}
		})
	}
}
