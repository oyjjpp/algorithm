package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{-1, head}
	var getKthFromEnd func(head *ListNode, k int) *ListNode
	getKthFromEnd = func(head *ListNode, k int) *ListNode {
		p1 := head
		// 先让P1走K步
		for i := 0; i < k; i++ {
			p1 = p1.Next
		}
		p2 := head
		// p1 和 p2 同时走n-k步
		for p1 != nil {
			p1 = p1.Next
			p2 = p2.Next
		}
		return p2
	}
	x := getKthFromEnd(dummy, n+1)
	x.Next = x.Next.Next
	return dummy.Next
}
