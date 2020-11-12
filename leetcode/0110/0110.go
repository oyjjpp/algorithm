package leetcode

import(
    "log"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isBalanced
// 平衡二叉树
// 前序遍历方式
func isBalanced(root *TreeNode) bool {
    // 如果节点为nil 则为平衡二叉树
    if root == nil {
        return true
    }
    
    // 左右子数的高度差
    log.Println(root.Val)
    log.Println("左子数的高度", maxDepth(root.Left))
    log.Println("右子数的高度", maxDepth(root.Right))
    if abs(maxDepth(root.Left)-maxDepth(root.Right)) > 1 {
        return false
    }
    
    if !isBalanced(root.Left){
        return false
    }
    return isBalanced(root.Right)
}

// maxDepth
// 二叉树高度
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

// max
// 取两个数的最大值
func max(left, right int) int {
	if left > right {
		return left
	}
	return right
}


// isBalancedV2
// 平衡二叉树
// 后续遍历
func isBalancedV2(root *TreeNode) bool {
    return height(root) >= 0
}

// height
func height(root *TreeNode) int {
    if root == nil {
        return 0
    }
    // 左右子数高度
    leftHeight := height(root.Left)
    rightHeight := height(root.Right)
    log.Println(root.Val, leftHeight, rightHeight)
    
    // 校验左节点和右节点高度是否为-1
    // 校验左右节点高度差是否大于1
    if leftHeight == -1 || rightHeight == -1 || abs(leftHeight - rightHeight) > 1 {
        return -1
    }
    
    // 求当前节点的高度 
    return max(leftHeight, rightHeight) + 1
}
