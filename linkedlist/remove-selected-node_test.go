package linkedlist

import (
	"reflect"
	"testing"
)

func TestRemoveSelectedNode(t *testing.T) {

	l := NewSentinelLinkedList()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)

	t.Run("1->2->3->4,给定节点2,删除节点2", func(t *testing.T) {
		RemoveSelectedNode(l, l.Head.Next.Next)
		want := NewSentinelLinkedList()
		want.Add(1)
		want.Add(3)
		want.Add(4)
		if !reflect.DeepEqual(l.Head, want.Head) {
			t.Fatalf("got:%+v,want;%+v", l.Head, want.Head)
		}
	})
}
