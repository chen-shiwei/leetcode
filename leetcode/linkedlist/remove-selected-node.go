package linkedlist

func RemoveSelectedNode(list *SentinelLinkedList, delete *Node) {
	if delete.Next == nil {
		tmp := list.Head
		for tmp.Next != delete {
			tmp = tmp.Next
		}
		tmp.Next = nil
	} else {
		// 直接修改值
		nextNode := delete.Next
		delete.data = delete.Next.data
		delete.Next = delete.Next.Next
		nextNode.Next = nil
	}
}
