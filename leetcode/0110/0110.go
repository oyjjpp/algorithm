package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isBalanced
// 平衡二叉树
func isBalanced(root *TreeNode) bool {
    if root == nil{
        return true
    }
    
    // 左右子数的告诉差
    if abs(height(root.Left)-height(root.Right)>1) {
        return false
    }
    
    if !isBalanced(root.Left){
        return false
    }
    return isBalanced(root.Right)
}

// height
// 求树的高度
func height(root *TreeNode) int {
    if root == nil {
        return 0
    }
    leftH := height(root.Left)
    rightH := height(root.Right)
    return max(leftH, rightH) + 1
}

// abs
// 求一个数的绝对值
func abs(num int) int {
    if num >0 {
        return num
    }
    return -num
}

func max(left, right int) int {
    if left > right {
        return left
    }
    return right
}
