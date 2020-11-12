package leetcode

import (
	"log"
	"testing"
)

func TestIsUnivalTree(t *testing.T) {
	data := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}
	rs := isUnivalTree(data)
	log.Println(rs)
}
