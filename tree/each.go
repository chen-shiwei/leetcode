package tree

import (
	"log"
)

type Node struct {
	val         int
	Left, Right *Node
}

func PrevEach(root *Node) {
	if root == nil {
		return
	}
	log.Println(root.val)
	PrevEach(root.Left)
	PrevEach(root.Right)

	return
}
func InEach(root *Node) {
	if root == nil {
		return
	}
	InEach(root.Left)
	log.Println(root.val)
	InEach(root.Right)

	return
}
func BeforeEach(root *Node) {
	if root == nil {
		return
	}
	BeforeEach(root.Left)
	BeforeEach(root.Right)
	log.Println(root.val)

	return
}
