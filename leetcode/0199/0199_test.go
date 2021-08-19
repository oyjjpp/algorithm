package leetcode

import "testing"

func TestRightSideView(t *testing.T) {
	data := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	rs := rightSideView(data)
	t.Log(rs)
}
