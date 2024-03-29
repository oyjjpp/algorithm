package leetcode

import "testing"

func TestZigzagLevelOrderV2(t *testing.T) {
	data := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	rs := zigzagLevelOrder(data)
	t.Log(rs)
}
