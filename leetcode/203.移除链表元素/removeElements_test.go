package _03_移除链表元素

import (
	"reflect"
	"testing"
)

func Test_removeElements(t *testing.T) {
	type args struct {
		head *ListNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: `输入：head = [1,2,6,3,4,5,6], val = 6
输出：[1,2,3,4,5]`, args: args{
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 6,
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
				},
			},
			val: 6,
		}, want: &ListNode{
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
		}},
		{name: `输入：head = [], val = 1
输出：[]`, args: args{
			head: nil,
			val:  1,
		}, want: nil},
		{name: `输入：head = [7,7,7,7], val = 7
输出：[]`, args: args{
			head: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 7,
					Next: &ListNode{
						Val: 7,
						Next: &ListNode{
							Val:  7,
							Next: nil,
						},
					},
				},
			},
			val: 7,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := removeElementsSentinel(tt.args.head, tt.args.val); !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("removeElementsSentinel() = %v, want %v", got, tt.want)
			//}
			if got := removeElementsRecursion(tt.args.head, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeElementsRecursion() = %v, want %v", got, tt.want)
			}

		})
	}
}
