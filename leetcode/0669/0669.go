package leetcode

import (
	"log"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// trimBST
// 修剪二叉搜索树
// @param root 二叉树
// @param low 最小边界
// @param high 最大边界
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil{
		return root
	}
	// 当前节点大于最大边界
	if root.Val>high {
		trimBST(root.Left, low, high)
	}
	if root.Val<low {
		trimBST(root.Right, low, high)
	}
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	
	return root
}
