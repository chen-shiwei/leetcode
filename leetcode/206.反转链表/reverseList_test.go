package _06_反转链表

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: `1,2,3,4,5`, args: args{head: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
						},
					},
				},
			},
		}}, want: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  1,
							Next: nil,
						},
					},
				},
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printListNode(tt.args.head)
			got := reverseListWithRecursion(tt.args.head)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseListWithRecursion() = %v, want %v", got, tt.want)
			}
			printListNode(got)
		})
	}
}

func printListNode(node *ListNode) {
	if node != nil {
		fmt.Printf("%+v =>", node)
		node = node.Next
		printListNode(node)
	} else {
		fmt.Println()
		return
	}
}
