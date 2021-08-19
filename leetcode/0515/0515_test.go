package leetcode

import "testing"

func TestLargestValues(t *testing.T) {
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
	rs := largestValues(data)
	t.Log(rs)
}
