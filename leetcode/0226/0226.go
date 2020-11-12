package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// invertTree
// 翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
	var node *TreeNode

	var preOrder func(root *TreeNode)
	preOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		node = root.Left
		root.Left = root.Right
		root.Right = node

		preOrder(root.Left)
		preOrder(root.Right)
	}
	preOrder(root)
	return root
}
