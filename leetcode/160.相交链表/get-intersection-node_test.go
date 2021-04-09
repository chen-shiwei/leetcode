package _60_相交链表

import (
	"reflect"
	"testing"
)

var list = &ListNode{
	Val: 8,
	Next: &ListNode{
		Val: 4,
		Next: &ListNode{
			Val:  5,
			Next: nil,
		},
	},
}

func TestGetIntersectionNode(t *testing.T) {
	type args struct {
		headA *ListNode
		headB *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "", args: args{
			headA: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  1,
					Next: list,
				},
			},
			headB: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val:  1,
						Next: list,
					},
				},
			},
		}, want: list},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetIntersectionNodeWithPointer(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
