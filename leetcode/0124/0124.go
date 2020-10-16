package leetcode

import (
    "math"
)

// 二叉树中的最大路径和

type TreeNode struct{
    Val int
    Left *TreeNode
    Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
    maxSum := math.MinInt32
    maxPath(root, *maxSum)
    return maxSum
}

func maxPath(root *TreeNode, maxSum *int) int {
    if root == nil {
        return 0
    }

    left := max(0, maxPathSum(root.Left))
    right := max(0, maxPathSum(root.Right))
    *maxSum = max(*maxSum, (left+right+root.Val))
    return max(left, right) + root.Val
}

// max
func max(left, right int) int{
    if left > right {
        return left
    }
    return right
}
