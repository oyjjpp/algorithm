package binary

import (
	"container/list"
	"log"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历
func preOrder(node *TreeNode) {
	if node == nil {
		return
	}
	log.Println(node.Val)
	preOrder(node.Left)
	preOrder(node.Right)
}

// 中序遍历
func midOrder(node *TreeNode) {
	if node == nil {
		return
	}
	midOrder(node.Left)
	log.Println(node.Val)
	midOrder(node.Right)
}

// 后续遍历
func subOrder(node *TreeNode) {
	if node == nil {
		return
	}
	subOrder(node.Left)
	subOrder(node.Right)
	log.Println(node.Val)
}

// 广度优先遍历（BFS）
// 使用数组实现简单队列
func bfsOrder(node *TreeNode) {
	if node == nil {
		return
	}

	var queue []*TreeNode
	queue = append(queue, node)

	for len(queue) > 0 {
		curNode := queue[0]
		if curNode != nil {
			log.Println(curNode.Val)
		}

		if curNode.Left != nil {
			queue = append(queue, curNode.Left)
		}

		if curNode.Right != nil {
			queue = append(queue, curNode.Right)
		}

		queue = queue[1:]
	}
}

// 广度优先遍历（BFS）
// 使用链表实现简单队列
func bfsListOrder(node *TreeNode) {
	if node == nil {
		return
	}

	var queue list.List
	queue.PushBack(node)

	for queue.Len() > 0 {
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
		}

		if curNode.Left != nil {
			queue.PushBack(curNode.Left)
		}

		if curNode.Right != nil {
			queue.PushBack(curNode.Right)
		}
	}
}

func dfsOrder(node *TreeNode) {
	if node == nil {
		return
	}

}
