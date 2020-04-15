package search

import "testing"

func Test_dfs(t *testing.T) {
	type args struct {
		node *Node
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "dfs", args: args{node: &Node{
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
			dfs(tt.args.node)
		})
	}
}

func Test_dfsWithStack(t *testing.T) {
	type args struct {
		node *Node
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "dfs", args: args{node: &Node{
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
			dfsWithStack(tt.args.node)
		})
	}
}
