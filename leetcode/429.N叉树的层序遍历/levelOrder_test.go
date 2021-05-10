package _29_N叉树的层序遍历

import (
	"reflect"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: `输入：root = [1,null,3,2,4,null,5,6]
输出：[[1],[3,2,4],[5,6]]`, args: args{root: &Node{
			Val: 1,
			Children: []*Node{
				&Node{
					Val: 3,
					Children: []*Node{
						&Node{
							Val:      5,
							Children: nil,
						},
						&Node{
							Val:      6,
							Children: nil,
						},
					},
				},
				&Node{
					Val:      2,
					Children: nil,
				},
				&Node{
					Val:      4,
					Children: nil,
				},
			},
		}}, want: [][]int{
			{1},
			{3, 2, 4},
			{5, 6},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
