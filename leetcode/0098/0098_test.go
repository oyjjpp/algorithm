package leetcode

import (
	"log"
	"testing"
)

// [5,1,4,null,null,3,6]
func TestIsValidBST(t *testing.T) {

	// data := &TreeNode{
	// 	Val:  5,
	// 	Left: &TreeNode{Val: 1},
	// 	Right: &TreeNode{
	// 		Val:   4,
	// 		Left:  &TreeNode{Val: 3},
	// 		Right: &TreeNode{Val: 6},
	// 	},
	// }

	/*
		data := &TreeNode{
			Val:2,
			Left:&TreeNode{Val:1},
			Right:&TreeNode{Val:3},
		}
	*/

	/*
		data := &TreeNode{
			Val:1,
			Left:&TreeNode{Val:1},
		}
	*/

	data := &TreeNode{
		Val:  2,
		Left: &TreeNode{Val: 1},
		Right: &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 6},
		},
	}
	rs := isValidBST(data)
	log.Println(rs)
}
