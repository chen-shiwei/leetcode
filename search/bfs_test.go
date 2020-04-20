package search

import "testing"

func Test_bfs(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "dfs", args: args{root: &Node{
			data: 1,
			LeftNode: &Node{
				data: 2,
				LeftNode: &Node{
					data:      5,
					LeftNode:  nil,
					MidNode:   nil,
					RightNode: nil,
				},
				RightNode: &Node{
					data:      6,
					LeftNode:  nil,
					MidNode:   nil,
					RightNode: nil,
				},
			},
			RightNode: &Node{
				data:     4,
				LeftNode: nil,
			},
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bfs(tt.args.root)
		})
	}
}
