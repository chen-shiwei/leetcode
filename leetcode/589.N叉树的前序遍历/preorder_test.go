package _89_N叉树的前序遍历

import (
	"reflect"
	"testing"
)

func Test_preorder(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: `输入：root = [1,null,3,2,4,null,5,6]
输出：[1,3,5,6,2,4]`, args: args{root: &Node{
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
		}}, want: []int{1, 3, 5, 6, 2, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
