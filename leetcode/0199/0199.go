package leetcode

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	data := make([]int, 0)

	var queue list.List
	queue.PushBack(root)

	for queue.Len() > 0 {
		number := queue.Len()

		for i := 0; i < number; i++ {
			element := queue.Front()
			queue.Remove(element)

			if element == nil {
				continue
			}

			node, ok := element.Value.(*TreeNode)
			if !ok {
				continue
			}

			if i == number-1 {
				data = append(data, node.Val)
			}

			if node.Left != nil {
				queue.PushBack(node.Left)
			}

			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}

	return data
}
