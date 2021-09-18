package leetcode

import (
	"container/list"
	"log"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 存储数据时候反序
func zigzagLevelOrder(node *TreeNode) [][]int {
	if node == nil {
		return nil
	}

	data := make([][]int, 0)

	var queue list.List
	queue.PushBack(node)

	evenNumbers := false

	for queue.Len() > 0 {

		// 分层次
		number := queue.Len()
		item := make([]int, number)
		for i := 0; i < number; i++ {

			index := i
			if evenNumbers {
				index = number - i - 1
			}

			// 读取一个元素
			element := queue.Front()
			queue.Remove(element)
			if element == nil {
				continue
			}

			curNode, ok := element.Value.(*TreeNode)
			if !ok {
				continue
			}

			if curNode != nil {
				log.Println(curNode.Val)
				item[index] = curNode.Val
			}

			if curNode.Left != nil {
				queue.PushBack(curNode.Left)
			}

			if curNode.Right != nil {
				queue.PushBack(curNode.Right)
			}
		}

		data = append(data, item)
		evenNumbers = !evenNumbers
	}

	return data
}
