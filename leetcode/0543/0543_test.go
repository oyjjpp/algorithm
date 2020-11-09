package leetcode

import (
	"testing"
	"log"
)

func TestDiameterOfBinaryTree(t *testing.T){
	data := &TreeNode{
		Val:1,
		Left:&TreeNode{
			Val:2,
			Left:&TreeNode{Val:4},
			Right:&TreeNode{Val:5},
		},
		Right:&TreeNode{Val:3},
	}
	
	rs := diameterOfBinaryTree(data)
	log.Println(rs)
}
