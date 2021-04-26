package tree

import (
	"fmt"
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
func preorderTraversal(root *Node) {
	var (
		stack  []*Node
		result []int
	)

	for root != nil || len(stack) != 0 {
		for root != nil {
			result = append(result, root.val)

			stack = append(stack, root)
			root = root.Left
		}

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}

	fmt.Println(result)
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

func inorderTraversal(root *Node) {
	var (
		stack  []*Node
		result []int
	)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.val)
		root = node.Right
	}

	fmt.Println(result)
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

func postorderTraversal(root *Node) {
	var (
		stack  []*Node
		result []int
	)
	var lastNode *Node
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		node := stack[len(stack)-1]
		if node.Right == nil || node.Right == lastNode {
			stack = stack[:len(stack)-1]
			result = append(result, node.val)
			lastNode = node
		} else {
			root = node.Right
		}
	}

	fmt.Println(result)
}
