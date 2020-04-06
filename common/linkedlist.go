package common

import "fmt"

// 138、复制带随机指针的链表
// 141、环形链表
// 148、排序链表
// 160、相交链表
// 206、反转链表
// 234、回文链表
// 237、删除链表中的节点
// 328、奇偶链表

type NodeList struct {
	Val    int
	Next   *Node
	Random *Node
}

// copyRandomList
// 复制带随机指针的链表
func copyRandomList(head *NodeList) *NodeList {
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseList
// 反转链表
// 迭代方法
func reverseList(head *ListNode) *ListNode {
	// 前指针节点
	var prev *ListNode
	cur := head

	for cur != nil {
		// 临时节点，暂时存放当前节点的下一个节点，用于后移
		temp := cur.Next
		// 将新指针转义到当前节点的后面
		cur.Next = prev
		// 更新最新的指针
		prev = cur
		// 更新原有指针
		cur = temp
	}
	return prev
}

// reverseList
// 反转链表
// 递归
func reverseListV2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 递归处理
	prev := reverseListV2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return prev
}

// reverseList
// 反转链表
// 尾递归
func reverseListV3(head *ListNode) *ListNode {
	return reverse(nil, head)
}
func reverse(pre, cur *ListNode) *ListNode {
	if cur == nil {
		return pre
	}
	next := cur.Next
	cur.Next = pre
	return reverse(cur, next)
}

// isPalindrome
// 是否为回文链表
// 思路
// 1、慢指针 指向头部
// 2、快指针 指向第二个节点
// 3、循环遍历链表，快指针每次走两个节点，慢指针每次走一个节点，当快指针遍历完链表，慢指针指向的位置就是中心点，
// 因为快指针从第二个节点开始走，所以不需要考虑奇偶的情况
// 4、当找到中心点后，把慢指针走过的节点从头到中心点截断，为待比较的第一部分链表
// 5、将从中心点到链表结尾的部分所有节点进行倒序操作，也就是链表的倒序。作为待比较的第二部分链表
// 6、比较两个链表每个节点是否相等，如果相等则为回文
func isPalindromeList(head *ListNode) bool {
	// 链表为空或者只有一个元素时则为回文链表
	if head == nil || head.Next == nil {
		return true
	}
	// 设置快指针,快指针从第二个开始，主要解决奇偶问题
	fast := head.Next
	// 慢指针
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	// 保留后面未遍历的链表
	curPre := slow.Next

	// 截断慢指针链表，仅保留前半部分
	slow.Next = nil
	var reverseV2 func(list *ListNode) *ListNode
	// 倒序遍历函数
	reverseV2 = func(list *ListNode) *ListNode {
		if list == nil || list.Next == nil {
			return list
		}
		pre := reverseV2(list.Next)
		list.Next.Next = list
		list.Next = nil
		return pre
	}
	// 将后面的链表倒序
	rightPre := reverseV2(curPre)
	for rightPre != nil && head != nil {
		// 比较两个值
		if rightPre.Val != head.Val {
			return false
		}
		rightPre = rightPre.Next
		head = head.Next
	}
	return true
}

func PrintList(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}

// deleteNode
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

// getIntersectionNode
// 相交链表
// 暴力破解【提交超时】
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	cur := headA
	for cur != nil {
		pre := headB
		for pre != nil {
			if cur.Val == pre.Val {
				return pre
			}
		}
	}
	return nil
}

// getIntersectionNode
// 相交链表
// 双指针方式
func getIntersectionNodeV2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	curA := headA
	curB := headB
	for curA != curB {
		if curA == nil {
			curA = headB
		} else {
			curA = curA.Next
		}

		if curB == nil {
			curB = headA
		} else {
			curB = curB.Next
		}
	}
	return curA
}

// getIntersectionNode
// 相交链表
// hash 方式
func getIntersectionNodeV3(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	hashMap := map[*ListNode]bool{}
	curA := headA
	curB := headB
	for curA != nil {
		hashMap[curA] = true
		curA = curA.Next
	}
	for curB != nil {
		if hashMap[curB] {
			return curB
		}
		curB = curB.Next
	}
	return nil
}

// oddEvenList
// 奇偶链表
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	// 存放奇数链表
	oneList := head
	// 存放偶数链表,保持链表的头
	twoList := head.Next
	curPre := twoList

	for curPre != nil && curPre.Next != nil {
		// 奇数位置
		oneList.Next = curPre.Next
		oneList = oneList.Next

		// 偶数位置
		curPre.Next = oneList.Next
		curPre = curPre.Next
	}
	// 将偶数链表头连接到奇数链表的尾
	oneList.Next = twoList
	return head
}
