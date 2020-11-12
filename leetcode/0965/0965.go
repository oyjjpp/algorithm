package leetcode

import (
	"log"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isUnivalTree
// 单值二叉树
func isUnivalTree(root *TreeNode) bool {
	data := map[int]bool{}

	var preOrder func(root *TreeNode)
	preOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		if _, ok := data[root.Val]; !ok {
			data[root.Val] = true
		}
		preOrder(root.Left)
		preOrder(root.Right)
	}

	preOrder(root)
	log.Println(data)
	return len(data) == 1
}
