package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// trimBST
// 修剪二叉搜索树
// @param root 二叉树
// @param low 最小边界
// @param high 最大边界
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return root
	}
	// 当前节点大于最大边界，则当前节点及右节点已经越界，所以只保留左子数
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}
	// 当前节点小于最小边界，则当前节点及左节点已经越界，所以值保留右子数
	if root.Val < low {
		return trimBST(root.Right, low, high)
	}
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)

	return root
}
