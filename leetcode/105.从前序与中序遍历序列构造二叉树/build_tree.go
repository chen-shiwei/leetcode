package _05_从前序与中序遍历序列构造二叉树

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

func buildTree(preorder []int, inorder []int) *TreeNode {
	l := len(preorder)
	if l < 1 {
		return nil
	}
	if l != len(inorder) {
		return nil
	}

	rootVal := preorder[0]
	rootNode := &TreeNode{
		Val:   rootVal,
		Left:  nil,
		Right: nil,
	}
	preorder = preorder[1:]

	// 找到根节点在中序遍历的位置
	// 根据中序遍历的特性 root节点左边的都是root节点的左子树，数量为(rootIndex+1)
	var rootIndex int

	for i, v := range inorder {
		if v == rootVal {
			rootIndex = i
			break
		}
	}

	leftInorder := inorder[:rootIndex]
	rightInorder := inorder[rootIndex+1:]

	leftPreorder := preorder[:len(leftInorder)]
	rightPreorder := preorder[len(leftInorder):]

	// 前序遍历的根节点的后面是左子树，左子树的切片为(1:rootIndex+1)
	rootNode.Left = buildTree(leftPreorder, leftInorder)
	rootNode.Right = buildTree(rightPreorder, rightInorder)
	return rootNode
}
