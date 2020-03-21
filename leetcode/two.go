// 两数相加
/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
*/
package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addNode(head, node *ListNode) *ListNode {
	temp := head
	for {
		// 当前头结点为空
		if head == nil {
			head = node
			return head
		}

		// 找到最后一个节点
		if temp.Next == nil {
			break
		}
		temp = temp.Next
	}
	temp.Next = node
	return head
}

func show(head *ListNode) {
	temp := head
	for {
		if temp == nil {
			break
		}
		fmt.Printf("%d", temp.Val)
		temp = temp.Next
		if temp != nil {
			fmt.Printf("-->")
		}
	}
}

// listNodeToSlice
func listNodeToSlice(list *ListNode) []int {
	data := make([]int, 0)
	temp := list
	for {
		if temp == nil {
			break
		}
		data = append(data, temp.Val)
		temp = temp.Next
	}
	return data
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	data1 := listNodeToSlice(l1)
	data2 := listNodeToSlice(l2)
	maxData := make([]int, 0)
	minData := make([]int, 0)
	if len(data1) <= len(data2) {
		maxData = data2
		minData = data1
	} else {
		maxData = data1
		minData = data2
	}
	maxLen := len(maxData)
	minLen := len(minData)
	rs := make([]int, maxLen+1)
	var head *ListNode
	for i, _ := range maxData {
		sum := 0
		if i <= (minLen - 1) {
			sum = minData[i] + maxData[i]
		} else {
			sum = maxData[i]
		}
		rs[i] = rs[i] + sum
		if rs[i] >= 10 {
			rs[i+1] = 1
			rs[i] = rs[i] - 10
		}
		head = addNode(head, &ListNode{
			Val:  rs[i],
			Next: nil,
		})
	}
	num := rs[len(rs)-1]
	if num != 0 {
		head = addNode(head, &ListNode{
			Val:  num,
			Next: nil,
		})
	}
	return head
}

func addTwoNumbersV2(l1, l2 *ListNode) *ListNode {
	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	curNode := head

	// 存上一位的和并满10进1
	sum := 0
	one, two := l1, l2
	for one != nil || two != nil {
		x, y := 0, 0
		if one != nil {
			x = one.Val
		}
		if two != nil {
			y = two.Val
		}
		// 当前位的和
		total := sum + x + y
		// 大于10则进位
		sum = total / 10

		curNode.Next = &ListNode{
			Val:  total % 10,
			Next: nil,
		}
		curNode = curNode.Next
		if one != nil {
			one = one.Next
		}
		if two != nil {
			two = two.Next
		}
	}

	// 存在进位情况
	if sum >= 0 {
		curNode.Next = &ListNode{
			Val:  sum,
			Next: nil,
		}
	}
	return head.Next
}
