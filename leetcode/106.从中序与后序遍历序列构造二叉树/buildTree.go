package _06_从中序与后序遍历序列构造二叉树

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

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	// 从后序遍历获取root
	rootVal := postorder[len(postorder)-1]
	root := &TreeNode{
		Val: rootVal,
	}
	if len(postorder) == 1 {
		return root
	}
	// 从后序遍历删除root
	postorder = postorder[:len(postorder)-1]

	// 获取root在中序遍历的位置
	var rootIdx int
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == rootVal {
			rootIdx = i
			break
		}
	}

	// root在中序遍历的位置切分左右子树，后续根据前序切分的数组长度切分后续数组
	root.Left = buildTree(inorder[:rootIdx], postorder[:len(inorder[:rootIdx])])
	root.Right = buildTree(inorder[rootIdx+1:], postorder[len(inorder[:rootIdx]):])
	return root
}
