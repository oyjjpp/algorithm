package hot100

import "log"

func maxDepth(root *TreeNode) int {
	var dp func(root *TreeNode) int
	dp = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := dp(root.Left)
		right := dp(root.Right)

		return max(left, right) + 1
	}
	return dp(root)
}

func printBinaryLevel(root *TreeNode) {
	if root == nil {
		return
	}

	var dp func(root *TreeNode, level int)
	dp = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		log.Printf("node:%d,level:%d", root.Val, level)
		dp(root.Left, level+1)
		dp(root.Right, level+1)
	}
}
