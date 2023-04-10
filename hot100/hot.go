package hot100

import (
	"log"
)

// 1.两数之和
func twoSum(nums []int, target int) []int {
	// 借助map实现
	data := map[int]int{}
	for k, v := range nums {
		temp := target - v
		if index, ok := data[temp]; ok {
			return []int{index, k}
		}
		data[v] = k
	}
	return nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 2.两数相加
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

// 3.无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	window := map[byte]int{}
	left, right := 0, 0
	length := 0

	for right < len(s) {
		ele := s[right]
		right++
		window[ele]++

		// 缩小窗口
		for window[ele] > 1 {
			del_ele := s[left]
			left++
			window[del_ele]--
		}
		length = max(length, right-left)
	}
	return length
}

// 4. 寻找两个正序数组的中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	data := make([]int, 0)

	i, j := 0, 0

	for i < len(nums1) && j < len(nums2) {
		left := 0
		if i < len(nums1) {
			left = nums1[i]
			i++
		}

		right := 0
		if j < len(nums2) {
			right = nums2[j]
			j++
		}

		if left < right {
			data = append(data, left)
			j--
		} else {
			data = append(data, right)
			i--
		}
	}
	log.Println(i, j, nums1[i:])
	if i < len(nums1) {
		data = append(data, nums1[i:]...)
	}

	if j < len(nums2) {
		data = append(data, nums2[j:]...)
	}
	log.Println(data)
	index := len(data) / 2
	if len(data)%2 == 0 {
		return float64(data[index]+data[index-1]) / 2
	} else {
		return float64(data[index])
	}
}

// 4. 寻找两个正序数组的中位数
func findMedianSortedArraysV(nums1 []int, nums2 []int) float64 {
	// 假设第一个数组长度较短
	m, n := len(nums1), len(nums2)
	if m > n {
		nums1, nums2, m, n = nums2, nums1, n, m
	}

	// 初始位置、较短数组长度、中位数所在位置  halfLen==(m+n+1)/2 ???
	imin, imax, halfLen := 0, m, (m+n+1)/2

	for imin <= imax {
		// 短数组移动位置
		i := (imin + imax) / 2

		// 长数组移动位置
		j := halfLen - i

		if i < m && nums2[j-1] > nums1[i] {
			// num2部分大于num1
			imin = i + 1
		} else if i > 0 && nums1[i-1] > nums2[j] {
			// nums1部分大于num2
			imax = i - 1
		} else {
			maxOfLeft := 0
			if i == 0 {
				maxOfLeft = nums2[j-1]
			} else if j == 0 {
				maxOfLeft = nums1[i-1]
			} else {
				maxOfLeft = max(nums1[i-1], nums2[j-1])
			}

			if (m+n)%2 == 1 {
				return float64(maxOfLeft)
			}

			minOfRight := 0
			if i == m {
				minOfRight = nums2[j]
			} else if j == n {
				minOfRight = nums1[i]
			} else {
				minOfRight = min(nums1[i], nums2[j])
			}

			return float64(maxOfLeft+minOfRight) / 2
		}
	}
	return 0.0
}

// 76. 最小覆盖子串
func minWindow(s string, t string) string {
	need := map[byte]int{}

	for i := 0; i < len(t); i++ {
		cur := t[i]
		need[cur]++
	}

	window := map[byte]int{}
	valid := 0

	maxNum := 1<<31 - 1
	left, right := 0, 0

	start, length := 0, maxNum

	for right < len(s) {
		cur := s[right]
		right++

		if _, ok := need[cur]; ok {
			window[cur]++

			if window[cur] == need[cur] {
				valid++
			}
		}
		// A DOBECODEBA NC ABC BANC
		// 窗口收缩
		for valid == len(need) {
			if (right - left) < length {
				start = left
				length = right - left
			}

			delELe := s[left]
			left++
			// ??????????
			if _, ok := need[delELe]; ok {
				if window[delELe] == need[delELe] {
					valid--
				}
				window[delELe]--
			}
		}

	}
	if length == maxNum {
		return ""
	}

	return s[start : start+length]
}
