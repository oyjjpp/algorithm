package leetcode

import "testing"

func TestMaxPathSum(t *testing.T) {
	data := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	data1 := &TreeNode{
		Val:  -10,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	rs := maxPathSum(data)
	t.Log(rs)
	rs1 := maxPathSum(data1)
	t.Log(rs1)
}
