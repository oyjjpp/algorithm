package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// kthLargest
// @desc 剑指 Offer 54. 二叉搜索树的第k大节点
func kthLargest(root *TreeNode, k int) int {
	data := []int{}
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		data = append(data, root.Val)
		inorder(root.Right)
	}

	inorder(root)
	return data[len(data)-k]
}

// kthLargest
// @desc 剑指 Offer 54. 二叉搜索树的第k大节点
func kthLargestV2(root *TreeNode, k int) int {
	var rs int
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Right)
		k--
		if k == 0 {
			rs = root.Val
			return
		}
		inorder(root.Left)
	}

	inorder(root)
	return rs
}
