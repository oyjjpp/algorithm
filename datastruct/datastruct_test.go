package datastruct

import "testing"

func TestPreorderTraversal(t *testing.T) {
	data := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 10}},
		Right: &TreeNode{Val: 1},
	}
	rs := postorderTraversal(data)
	t.Log(rs)
}
