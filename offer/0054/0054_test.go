package leetcode

import(
	"testing"
	"log"
)

func TestKthLargest(t *testing.T){
	data := &TreeNode{
		Val:3,
		Left:&TreeNode{
			Val:1,
			Right:&TreeNode{
				Val:2,
			},
		},
		Right:&TreeNode{
			Val:4,
		},
	}
	rs := kthLargestV2(data, 1)
	log.Println(rs)
}
