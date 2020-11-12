package leetcode

import (
	"log"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	data := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	length := maxDepth(data)
	log.Println(length)
}
