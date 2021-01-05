package leetcode

import (
	"log"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isValidBST
// 是否有有效二叉树
func isValidBST(root *TreeNode) bool {
	var pre *TreeNode
	isValid := true

	// 通过中序遍历的有序性，校验是否为有效二叉树
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		if pre == nil {
			pre = root
		} else {
			if pre.Val >= root.Val {
				isValid = false
				return
			}
			pre = root
		}
		log.Println(pre.Val)
		inorder(root.Right)
	}
	inorder(root)
	return isValid
}
