// 二叉树中的最大路径和
package leetcode

import (
	"log"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := -1 << 31
	maxPath(root, &maxSum)
	return maxSum
}

func maxPath(root *TreeNode, maxSum *int) int {
	if root == nil {
		return 0
	}

	// 求左右子数最大路径和
	left := max(0, maxPath(root.Left, maxSum))
	right := max(0, maxPath(root.Right, maxSum))
	temp := *maxSum
	// 当前子数的路径最大和与最大值对比
	*maxSum = max(*maxSum, (left + right + root.Val))
	log.Println(root.Val, temp, *maxSum, (left + right + root.Val))
	// 为什么max(left, right):请最大路径root+left,root+right
	return max(left, right) + root.Val
}

// max
func max(left, right int) int {
	if left > right {
		return left
	}
	return right
}
