package hot100

import (
	"container/heap"
	"log"
)

/*
 双指针技巧秒杀七道链表题目
*/

// 21. 合并两个有序链表
// 双指针
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 使用虚拟头结点，避免解决data空指针问题
	head := &ListNode{}
	data := head
	node1 := list1
	node2 := list2

	for node1 != nil && node2 != nil {
		if node1.Val > node2.Val {
			data.Next = node2
			node2 = node2.Next
		} else {
			data.Next = node1
			node1 = node1.Next
		}
		data = data.Next
	}

	if node1 != nil {
		data.Next = node1
	}
	if node2 != nil {
		data.Next = node2
	}
	return head.Next
}

// 86. 分隔链表
// 双指针
func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	small := &ListNode{}
	smallTail := small

	dummy := &ListNode{0, head}

	pre := dummy
	cur := head

	for cur != nil {
		if cur.Val < x {
			smallTail.Next = cur
			smallTail = smallTail.Next
			//相当于删除结点，pre不用动
			pre.Next = cur.Next
			cur = cur.Next

		} else {
			//无事发生,一起移动
			pre = cur
			cur = cur.Next
		}

	}

	smallTail.Next = dummy.Next
	return small.Next
}

// 23. 合并 K 个升序链表
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	// 虚拟头结点
	dummy := &ListNode{}
	p := dummy
	log.Println(p)

	// 优先级队列，最小堆
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	// 将数据加入最小堆
	for _, head := range lists {
		if head != nil {
			heap.Push(&pq, head)
		}
	}

	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*ListNode)
		p.Next = node
		if node.Next != nil {
			heap.Push(&pq, node.Next)
		}
		p = p.Next
	}
	return dummy.Next
}

// 优先级队列（二叉堆）
type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	node := x.(*ListNode)
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[0 : n-1]
	return node
}

// 剑指 Offer 22. 链表中倒数第k个节点
func getKthFromEnd(head *ListNode, k int) *ListNode {
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

// 19. 删除链表的倒数第 N 个结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 使用虚拟节点 避免越界
	dummy := &ListNode{-1, head}

	// 搜索倒数第N个节点
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

// 876. 链表的中间结点
func middleNode(head *ListNode) *ListNode {
	left, right := head, head

	for right != nil && right.Next != nil {
		left = left.Next
		right = right.Next.Next
	}
	return left
}

// 141. 环形链表
func hasCycle(head *ListNode) bool {
	left, right := head, head

	for right != nil && right.Next != nil {
		left = left.Next
		right = right.Next.Next

		if left == right {
			return true
		}
	}
	return false
}

// 剑指 Offer II 022. 链表中环的入口节点
func detectCycle(head *ListNode) *ListNode {
	left, right := head, head

	// 通过快慢指针寻找到环的入口
	for right != nil && right.Next != nil {
		left = left.Next
		right = right.Next.Next

		if left == right {
			break
		}
	}
	if right == nil || right.Next == nil {
		return nil
	}
	left = head

	for left != right {
		left = left.Next
		right = right.Next
	}
	return left
}

// 160. 相交链表
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	data := map[*ListNode]struct{}{}
	left := headA
	for left != nil {
		data[left] = struct{}{}
		left = left.Next
	}

	right := headB
	for right != nil {
		if _, ok := data[right]; ok {
			return right
		}
		right = right.Next
	}

	return nil
}

// 160. 相交链表
func getIntersectionNodeV2(headA, headB *ListNode) *ListNode {
	p1, p2 := headA, headB
	for p1 != p2 {
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}

		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	return p1
}
