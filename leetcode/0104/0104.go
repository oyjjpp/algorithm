package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// maxDepth
// 二叉树深度
// DFS(深度优先搜素)
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// max
// 求两个点最大的值
func max(left, right int) int {
	if left > right {
		return left
	}
	return right
}
