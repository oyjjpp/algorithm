package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// sortedArrayToBST
// 数组转换为二叉搜索树
// 使用中序遍历的方式
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	var preOrder func(nums []int, left, right int) *TreeNode
	preOrder = func(nums []int, left, right int) *TreeNode {
		if right < left {
			return nil
		}
		// 去中间元素做为根节点，达到二叉树高度平衡
		mid := (left + right) / 2
		root := &TreeNode{Val: nums[mid]}

		// 以左右分割达到二叉搜索树中序遍历结果升序的结果
		root.Left = preOrder(nums, left, mid-1)
		root.Right = preOrder(nums, mid+1, right)
		return root
	}

	return preOrder(nums, 0, len(nums)-1)
}
