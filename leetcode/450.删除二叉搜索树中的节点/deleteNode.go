package _50_删除二叉搜索树中的节点

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val == key {
		// 要删除的节点等于叶子节点
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			node := root.Right
			for node.Left != nil {
				node = node.Left
			}

			node.Left = root.Left
			return root.Right
		}

	}

	return root
}

func successor(node *TreeNode) int {
	node = node.Right
	for node.Left != nil {
		node = node.Left
	}
	return node.Val
}
func predecessor(node *TreeNode) int {
	node = node.Left
	for node.Right != nil {
		node = node.Right
	}
	return node.Val
}
