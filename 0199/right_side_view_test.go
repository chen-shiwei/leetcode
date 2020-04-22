package _0199

import (
	"reflect"
	"testing"
)

func Test_rightSideView(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: `输入: [1,2,3,null,5,null,4]
				输出: [1, 3, 4]
				解释:
				
				   1            <---
				 /   \
				2     3         <---
				 \     \
				  5     4       <---`, args: args{root: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:  2,
				Left: nil,
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
		}}, want: []int{1, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name+"bfs", func(t *testing.T) {
			if got := rightSideView(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rightSideView() = %v, want %v", got, tt.want)
			}

		})
		t.Run(tt.name+"dfs", func(t *testing.T) {
			if got := rightSideViewDfs(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rightSideView() = %v, want %v", got, tt.want)
			}

		})
	}
}
