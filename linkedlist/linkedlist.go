package linkedlist

type Node struct {
	data int
	Next *Node
}

type LinkedList struct {
	length int
	Head   *Node
}

func (l *LinkedList) Add(n int) {
	newNode := &Node{
		data: n,
		Next: nil,
	}
	if l.Head == nil {
		l.Head = newNode
	} else {
		tmp := l.Head
		for tmp.Next != nil {
			tmp = tmp.Next
		}
		tmp.Next = newNode
	}
	l.length++
}

type SentinelLinkedList struct {
	LinkedList
}

func NewSentinelLinkedList() *SentinelLinkedList {
	return &SentinelLinkedList{
		LinkedList{
			length: 0,
			Head: &Node{
				data: 0,
				Next: nil,
			},
		},
	}
}

func (l *SentinelLinkedList) Add(n int) {
	tmp := l.Head
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = &Node{
		data: n,
		Next: nil,
	}
	l.length++
}

func (l *SentinelLinkedList) HeadInsert(n int) {
	newNode := &Node{
		data: n,
		Next: nil,
	}

	newNode.Next = l.Head.Next
	l.Head.Next = newNode
	l.length++
}

func (l *SentinelLinkedList) Create(args []int) {
	for _, n := range args {
		l.HeadInsert(n)
	}
}
