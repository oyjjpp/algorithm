package hot100

import "log"

// 9 9 9 9 9 9 9
// 9 9 9 9
// 8 9 9 9 0 0 0 1

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{
		Val:  0,
		Next: nil,
	}

	data := head
	sum := 0
	node1 := l1
	node2 := l2

	for node1 != nil || node2 != nil {
		val1 := 0
		if node1 != nil {
			val1 = node1.Val
			node1 = node1.Next
		}

		val2 := 0
		if node2 != nil {
			val2 = node2.Val
			node2 = node2.Next
		}

		total := val1 + val2 + sum

		// 考虑进位
		data.Next = &ListNode{
			Val:  total % 10,
			Next: nil,
		}
		sum = total / 10
		data = data.Next
	}
	if sum > 0 {
		data.Next = &ListNode{
			Val:  sum,
			Next: nil,
		}
	}
	return head.Next
}

func scanList(node *ListNode) {
	head := node
	for head != nil {
		log.Println(head.Val)
		head = head.Next
	}
}
