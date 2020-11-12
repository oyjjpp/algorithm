package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isSymmetric
// 对称二叉树
func isSymmetric(root *TreeNode) bool {
	// 根节点为nil则是对称二叉树
	if root == nil {
		return true
	}
	var order func(nodeL, nodeR *TreeNode) bool
	order = func(nodeL, nodeR *TreeNode) bool {
		// 同时为空
		if nodeL == nil && nodeR == nil {
			return true
		}
		// 一个为空一个不为空
		if nodeL == nil && nodeR != nil {
			return false
		}
		if nodeL != nil && nodeR == nil {
			return false
		}
		// 递归校验 左节点的左子数与右节点的右子数 左节点的右子数和右节点的左子数
		if nodeL.Val == nodeR.Val {
			return order(nodeL.Left, nodeR.Right) && order(nodeL.Right, nodeR.Left)
		}
		return false
	}
	return order(root.Left, root.Right)
}
