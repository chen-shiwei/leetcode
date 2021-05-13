package _35_二叉搜索树的最近公共祖先

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	if root.Val > p.Val && root.Val > q.Val {
		left := lowestCommonAncestor(root.Left, p, q)
		if left != nil {
			return left
		}
	}
	if root.Val < p.Val && root.Val < q.Val {
		right := lowestCommonAncestor(root.Right, p, q)
		if right != nil {
			return right
		}
	}

	return root
}
