package leetcode

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}

	data := make([]float64, 0)

	var queue list.List
	queue.PushBack(root)

	for queue.Len() > 0 {
		number := queue.Len()
		total := 0
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

			total += node.Val

			if node.Left != nil {
				queue.PushBack(node.Left)
			}

			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		data = append(data, float64(total)/float64(number))
	}

	return data
}
