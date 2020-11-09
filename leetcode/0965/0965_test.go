package leetcode

import(
	"testing"
	"log"
)

func TestIsUnivalTree(t *testing.T){
	data := &TreeNode{
		Val:1,
		Left:&TreeNode{
			Val:1,
		},
		Right:&TreeNode{
			Val:2,
		},
	}
	rs := isUnivalTree(data)
	log.Println(rs)
}
