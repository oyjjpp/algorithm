package leetcode

// 二叉树中的最大路径和

type TreeNode struct{
    Val int
    Left *TreeNode
    Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
    if root == nil {
        return 0
    }

    left := max(0, maxPathSum(root.Left))
    right := max(0, maxPathSum(root.Right))

}

// max
func max(left, right int) int{
    if left > right {
        return left
    }
    return right
}
