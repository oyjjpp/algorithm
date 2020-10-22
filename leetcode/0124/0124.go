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
/*
data1 := &TreeNode{
	Val:  -10,
	Left: &TreeNode{Val: 9},
	Right: &TreeNode{
		Val:   20,
		Left:  &TreeNode{Val: 15},
		Right: &TreeNode{Val: 7},
	},
}
*/
func maxPath(root *TreeNode, maxSum *int) int {
	if root == nil {
		return 0
	}

	left := max(0, maxPath(root.Left, maxSum))
	right := max(0, maxPath(root.Right, maxSum))
	temp := *maxSum
	*maxSum = max(*maxSum, (left + right + root.Val))
	log.Println(root.Val, temp, *maxSum, (left + right + root.Val))
	return max(left, right) + root.Val
}

// max
func max(left, right int) int {
	if left > right {
		return left
	}
	return right
}
