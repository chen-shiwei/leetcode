package tree

import (
	"testing"
)

func TestPrevEach(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "前序遍历", args: args{root: &Node{
			val: 1,
			Left: &Node{
				val: 2,
				Left: &Node{
					val:   3,
					Left:  nil,
					Right: nil,
				},
				Right: &Node{
					val:   4,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &Node{
				val:  5,
				Left: nil,
				Right: &Node{
					val:   6,
					Left:  nil,
					Right: nil,
				},
			},
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrevEach(tt.args.root)
		})
	}
}
func TestInEach(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "中序遍历", args: args{root: &Node{
			val: 1,
			Left: &Node{
				val: 2,
				Left: &Node{
					val:   3,
					Left:  nil,
					Right: nil,
				},
				Right: &Node{
					val:   4,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &Node{
				val:  5,
				Left: nil,
				Right: &Node{
					val:   6,
					Left:  nil,
					Right: nil,
				},
			},
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InEach(tt.args.root)
		})
	}
}
func TestBeforeEach(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "后续遍历", args: args{root: &Node{
			val: 1,
			Left: &Node{
				val: 2,
				Left: &Node{
					val:   3,
					Left:  nil,
					Right: nil,
				},
				Right: &Node{
					val:   4,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &Node{
				val:  5,
				Left: nil,
				Right: &Node{
					val:   6,
					Left:  nil,
					Right: nil,
				},
			},
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BeforeEach(tt.args.root)
		})
	}
}
