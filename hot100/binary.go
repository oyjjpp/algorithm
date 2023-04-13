package hot100

import "log"

func maxDepth(root *TreeNode) int {
	var dp func(root *TreeNode) int
	dp = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := dp(root.Left)
		right := dp(root.Right)

		return max(left, right) + 1
	}
	return dp(root)
}

// 扫描二叉树节点所在层次
func printBinaryLevel(root *TreeNode) {
	if root == nil {
		return
	}

	var dp func(root *TreeNode, level int)
	dp = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		dp(root.Left, level+1)
		log.Printf("node:%d,level:%d", root.Val, level)
		dp(root.Right, level+1)

	}
	dp(root, 0)
}

// 543. 二叉树的直径
func diameterOfBinaryTree(root *TreeNode) int {
	var dp func(root *TreeNode) int

	// 可能是左右加起来最大 maxNumber
	maxNumber := 0
	dp = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := dp(root.Left)
		right := dp(root.Right)
		maxNumber = max(maxNumber, left+right)
		return max(left, right) + 1
	}

	dp(root)
	return maxNumber
}

// 515. 在每个树行中找最大值
func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	data := make([]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		sz := len(queue)

		maxNumber := -1 << 10
		for i := 0; i < sz; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Val > maxNumber {
				maxNumber = node.Val
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		data = append(data, maxNumber)
	}
	return data
}

// 决策树
// 树枝 节点
