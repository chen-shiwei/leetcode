package linkedlist

import (
	"reflect"
	"testing"
)

func TestLinkedList_Add(t *testing.T) {
	type fields struct {
		length int
		Head   *Node
	}
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "add",
			fields: fields{length: 0, Head: &Node{
				data: 1,
				Next: &Node{
					data: 2,
					Next: nil,
				},
			}},
			args: args{n: 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				length: tt.fields.length,
				Head:   tt.fields.Head,
			}
			l.Add(tt.args.n)

			want := &Node{
				data: 1,
				Next: &Node{
					data: 2,
					Next: &Node{
						data: 3,
						Next: nil,
					},
				},
			}
			if !reflect.DeepEqual(want, l.Head) {
				t.Fail()
			}
		})
	}
}

func TestSentinelLinkedList_Add(t *testing.T) {
	type fields struct {
		LinkedList LinkedList
	}
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "sentinel add", args: args{n: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewSentinelLinkedList()
			l.Add(tt.args.n)

			want := &Node{
				data: 0,
				Next: &Node{
					data: 1,
					Next: nil,
				},
			}
			if !reflect.DeepEqual(want, l.Head) {
				t.Fail()
			}
		})
	}
}

func TestSentinelLinkedList_Create(t *testing.T) {
	type fields struct {
		LinkedList LinkedList
	}
	type args struct {
		args []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Node
	}{
		{name: "头插法构造链表", args: args{args: []int{1, 2, 3, 4}}, want: &Node{
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
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewSentinelLinkedList()
			l.Create(tt.args.args)
			if !reflect.DeepEqual(l.Head.Next, tt.want) {
				t.Fail()
			}
		})
	}
}
