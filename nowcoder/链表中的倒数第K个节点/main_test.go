package 链表中的倒数第K个节点

import (
	"reflect"
	"testing"
)

func TestFindKthToTail(t *testing.T) {
	type args struct {
		pListHead *ListNode
		k         int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: ``, args: args{
			pListHead: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 6,
							Next: &ListNode{
								Val: 9,
								Next: &ListNode{
									Val: 11,
									Next: &ListNode{
										Val:  12,
										Next: nil,
									},
								},
							},
						},
					},
				},
			},
			k: 3,
		}, want: &ListNode{
			Val:  9,
			Next: nil,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindKthToTail(tt.args.pListHead, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindKthToTail() = %v, want %v", got, tt.want)
			}
		})
	}
}
