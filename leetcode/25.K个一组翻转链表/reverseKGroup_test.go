package _5_K个一组翻转链表

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_reverseLinkedList(t *testing.T) {
	type args struct {
		head *ListNode
		tail *ListNode
	}

	var listNode = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  6,
							Next: nil,
						},
					},
				},
			},
		},
	}

	var (
		// 2
		head = listNode.Next
		tail = listNode.Next.Next.Next

		next = listNode.Next.Next.Next.Next
	)

	tests := []struct {
		name  string
		args  args
		want  *ListNode
		want1 *ListNode
	}{
		{name: `123456`, args: args{
			head: head,
			tail: tail,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := reverseLinkedList(tt.args.head, tt.args.tail)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseLinkedList() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("reverseLinkedList() got1 = %v, want %v", got1, tt.want1)
			}

			listNode.Next = got
			fmt.Println(next, next.Next)
			got1.Next = next

			var printList func(node *ListNode)
			printList = func(node *ListNode) {
				if node != nil {
					println(node.Val)
					node = node.Next
					printList(node)
					return
				}
				return
			}
			printList(listNode)
		})
	}
}

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: ``, args: args{
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
								Next: &ListNode{
									Val: 6,
									Next: &ListNode{
										Val:  7,
										Next: nil,
									},
								},
							},
						},
					},
				},
			},
			k: 2,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseKGroup(tt.args.head, tt.args.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
			var printList func(node *ListNode)
			printList = func(node *ListNode) {
				if node != nil {
					println(node.Val)
					node = node.Next
					printList(node)
					return
				}
				return
			}
			printList(got)

		})
	}
}
