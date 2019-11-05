package _019

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd1(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "1->2->3->4->5,2", args: args{
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
			n: 2,
		}, want: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		}},
		{name: "1,1", args: args{
			head: &ListNode{
				Val:  1,
				Next: nil,
			},
			n: 1,
		}, want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RemoveNthFromEnd1(tt.args.head, tt.args.n)
			fmt.Println("got")
			printListNode(got)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveNthFromEnd2(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "1->2->3->4->5,2",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val:  5,
									Next: nil,
								},
							},
						},
					},
				},
				n: 2,
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  5,
							Next: nil,
						},
					},
				},
			},
		},
		{name: "1,1", args: args{
			head: &ListNode{
				Val:  1,
				Next: nil,
			},
			n: 1,
		}, want: nil,
		},
		{name: "1,2 => 1", args: args{
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  2,
					Next: nil,
				},
			},
			n: 1,
		}, want: &ListNode{
			Val:  1,
			Next: nil,
		},
		},
		{name: "1,2 => 2", args: args{
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  2,
					Next: nil,
				},
			},
			n: 2,
		}, want: &ListNode{
			Val:  2,
			Next: nil,
		},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RemoveNthFromEnd2(tt.args.head, tt.args.n)
			fmt.Println("got")
			printListNode(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
