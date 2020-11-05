package leetcode

import(
	"testing"
	"log"
)

// [5,1,4,null,null,3,6]
func TestIsValidBST(t *testing.T){
	data := &TreeNode{
		Val:5,
		Left:&TreeNode{Val:1},
		Right:&TreeNode{
			Val:4,
			Left:&TreeNode{Val:3},
			Right:&TreeNode{Val:6},
		},
	}
	rs := isValidBST(data)
	log.Println(rs)
}
