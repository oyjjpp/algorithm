package leetcode

import (
	"log"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// addTwoNumbers
// 两数相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	data := head
	node1 := l1
	node2 := l2
	// 存储进位
	sum := 0
	for node1 != nil || node2 != nil {
		num1 := 0
		if node1 != nil {
			num1 = node1.Val
			node1 = node1.Next
		}

		num2 := 0
		if node2 != nil {
			num2 = node2.Val
			node2 = node2.Next
		}
		total := num1 + num2 + sum
		sum = total / 10
		total = total % 10
		log.Println("total", total)

		data.Next = &ListNode{
			Val:  total,
			Next: nil,
		}
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
