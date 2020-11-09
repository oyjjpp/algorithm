package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// diameterOfBinaryTree
// 二叉树的直径
func diameterOfBinaryTree(root *TreeNode) int {
	var maxPath int

	// 后续遍历二叉树
	var postOrder func(root *TreeNode) int
	postOrder = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		// 左右子数的最大节点数
		leftLen := postOrder(root.Left)
		rightLen := postOrder(root.Right)
		// 当前节点最大节点数与当前最大节点树对别
		maxPath = max(maxPath, leftLen+rightLen+1)
		return max(leftLen, rightLen) + 1
	}

	postOrder(root)
	// 一条路径的长度为该路径经过的节点减一
	if maxPath == 0 {
		return 0
	}
	return maxPath - 1
}

// max
// 求两个数的最大值
func max(left, right int) int {
	if left > right {
		return left
	}
	return right
}
