package _07_设计链表

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	linkedList := Constructor()
	linkedList.AddAtHead(1)
	printList(linkedList.head.Next)
	linkedList.AddAtTail(3)
	printList(linkedList.head.Next)
	linkedList.AddAtIndex(1, 2) //链表变为1-> 2-> 3
	printList(linkedList.head.Next)
	fmt.Println("返回2", linkedList.Get(1)) //返回2
	printList(linkedList.head.Next)
	linkedList.DeleteAtIndex(1) //现在链表是1-> 3
	printList(linkedList.head.Next)
	fmt.Println("返回3", linkedList.Get(1)) //返回3
}

func printList(node *ListNode) {
	if node == nil {
		fmt.Println()
		return
	}
	fmt.Printf("%+v =>", node)

	printList(node.Next)
}
