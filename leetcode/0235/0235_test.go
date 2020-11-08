package leetcode

import(
	"testing"
	"log"
	"encoding/json"
)

func TestLowestCommonAncestor(t *testing.T){
	data := &TreeNode{
		Val:6,
		Left:&TreeNode{
			Val:2,
			Left:&TreeNode{
				Val:0,
			},
			Right:&TreeNode{
				Val:4,
				Left:&TreeNode{Val:3},
				Right:&TreeNode{Val:5},
			},
		},
		Right:&TreeNode{
			Val:8,
			Left:&TreeNode{Val:7},
			Right:&TreeNode{Val:9},
		},
	}
	
	p := &TreeNode{Val:3}
	q := &TreeNode{Val:5}
	rs := lowestCommonAncestor(data, p, q)
	result, err := json.Marshal(rs)
	if err != nil {
		t.Error(err.Error())
	}
	log.Println(string(result))
}
