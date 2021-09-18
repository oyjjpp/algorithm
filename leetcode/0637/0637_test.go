package leetcode

import "testing"

func TestAverageOfLevels(t *testing.T) {
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
	rs := averageOfLevels(data)
	t.Log(rs)
}
