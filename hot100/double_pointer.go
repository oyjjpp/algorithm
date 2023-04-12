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
