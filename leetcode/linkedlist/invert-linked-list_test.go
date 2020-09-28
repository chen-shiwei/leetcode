package linkedlist

import (
	"reflect"
	"testing"
)

func Test_recursiveInvert(t *testing.T) {
	type args struct {
		node *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{name: "翻转链表", args: args{node: &Node{
			data: 4,
			Next: &Node{
				data: 3,
				Next: &Node{
					data: 2,
					Next: &Node{
						data: 1,
						Next: nil,
					},
				},
			},
		}}, want: &Node{
			data: 1,
			Next: &Node{
				data: 2,
				Next: &Node{
					data: 3,
					Next: &Node{
						data: 4,
						Next: nil,
					},
				},
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := recursiveInvert(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("recursiveInvert() = %v, want %v", got, tt.want)
			}

		})

	}
}

func Test_eachInvert(t *testing.T) {
	type args struct {
		node *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{name: "翻转链表", args: args{node: &Node{
			data: 0,
			Next: &Node{
				data: 4,
				Next: &Node{
					data: 3,
					Next: &Node{
						data: 2,
						Next: &Node{
							data: 1,
							Next: nil,
						},
					},
				},
			}}}, want: &Node{
			data: 0,
			Next: &Node{
				data: 1,
				Next: &Node{
					data: 2,
					Next: &Node{
						data: 3,
						Next: &Node{
							data: 4,
							Next: nil,
						},
					},
				},
			}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := eachInvert(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("recursiveInvert() = %v, want %v", got, tt.want)
			}
		})

	}
}
