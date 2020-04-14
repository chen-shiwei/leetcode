package _142

import (
	"reflect"
	"testing"
)

func Test_detectCycle(t *testing.T) {

	head := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val:  -4,
					Next: nil,
				},
			},
		},
	}
	head.Next.Next.Next.Next = head.Next

	t.Run(`输入：head = [3,2,0,-4], pos = 1
输出：tail connects to node index 1
解释：链表中有一个环，其尾部连接到第二个节点。`, func(t *testing.T) {
		if got := detectCycle(head); !reflect.DeepEqual(got, head.Next) {
			t.Errorf("detectCycle() = %v, want %v", got, head.Next)
		}
	})
}
