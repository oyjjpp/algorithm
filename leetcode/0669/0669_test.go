package leetcode

import (
	"encoding/json"
	"log"
	"testing"
)

func TestTrimBST(t *testing.T) {
	// TODO 为什么不是二叉搜索树？？？ 尴尬了 原数据不是二叉搜索树
	// data := &TreeNode{
	// 	Val:  5,
	// 	Left: &TreeNode{Val: 1},
	// 	Right: &TreeNode{
	// 		Val:   4,
	// 		Left:  &TreeNode{Val: 3},
	// 		Right: &TreeNode{Val: 6},
	// 	},
	// }
	// TODO 验证结果
	data := &TreeNode{
		Val:  2,
		Left: &TreeNode{Val: 1},
		Right: &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 6},
		},
	}
	inorder(data)
	rsData := trimBST(data, 3, 5)
	log.Println("修剪之后")
	
	inorder(data)
	rs, _ := json.Marshal(data)
	log.Println(string(rs))
	
	inorder(rsData)
	rs1, _ := json.Marshal(rsData)
	log.Println(string(rs1))
}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	log.Println(root.Val)
	inorder(root.Right)
}
