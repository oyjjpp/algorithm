package leetcode

import (
	"log"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// levelOrderBottom
// 二叉树的层次遍历 II
func levelOrderBottom(root *TreeNode) [][]int {
	data := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	log.Println(data)
	return nil
}
