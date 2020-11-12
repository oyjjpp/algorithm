package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isBalanced
// 平衡二叉树
func isBalanced(root *TreeNode) bool {
    // 如果节点为nil 则为平衡二叉树
    if root == nil {
        return true
    }
    
    // 左右子数的深度差
    if abs(maxDepth(root.Left)-maxDepth(root.Right)) > 1 {
        return false
    }
    
    if !isBalanced(root.Left){
        return false
    }
    return isBalanced(root.Right)
}

// maxDepth
// 二叉树深度
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// abs
// 求一个数的绝对值
func abs(num int) int {
	if num > 0 {
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


func isBalancedV2(root *TreeNode) bool {
    return height(root) >= 0
}

func height(root *TreeNode) int {
    if root == nil {
        return 0
    }
    leftHeight := height(root.Left)
    rightHeight := height(root.Right)
    if leftHeight == -1 || rightHeight == -1 || abs(leftHeight - rightHeight) > 1 {
        return -1
    }
    return max(leftHeight, rightHeight) + 1
}

