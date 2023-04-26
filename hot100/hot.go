package hot100

import (
	"log"
	"sort"
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

func twoSumV1(nums []int, target int) []int {
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	for left < right {
		mid := nums[left] + nums[right]
		if mid == target {
			return []int{left, right}
		} else if mid > target {
			right--
		} else if mid < target {
			left++
		}
	}
	return nil
}

// 15. 三数之和
/* 计算数组 nums 中所有和为 target 的三元组 */
func threeSumTarget(nums []int, target int) [][]int {
	var twoSumTarget func(nums []int, target int) [][]int
	twoSumTarget = func(nums []int, target int) [][]int {
		// nums 数组必须有序
		sort.Ints(nums)
		lo, hi := 0, len(nums)-1
		res := [][]int{}
		for lo < hi {
			sum := nums[lo] + nums[hi]
			left, right := nums[lo], nums[hi]
			if sum < target {
				for lo < hi && nums[lo] == left {
					lo++
				}
			} else if sum > target {
				for lo < hi && nums[hi] == right {
					hi--
				}
			} else {
				res = append(res, []int{left, right})
				for lo < hi && nums[lo] == left {
					lo++
				}
				for lo < hi && nums[hi] == right {
					hi--
				}
			}
		}
		return res
	}

	// 数组得排个序
	sort.Ints(nums)
	n := len(nums)
	res := [][]int{}
	// 穷举 threeSum 的第一个数
	for i := 0; i < n; i++ {
		// 对 target - nums[i] 计算 twoSum
		tuples := twoSumTarget(nums[i+1:], target-nums[i])
		// 如果存在满足条件的二元组，再加上 nums[i] 就是结果三元组
		for _, tuple := range tuples {
			tuple = append(tuple, nums[i])
			res = append(res, tuple)
		}
		// 跳过第一个数字重复的情况，否则会出现重复结果
		for i < n-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}

// 18. 四数之和
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	length := len(nums)
	for i := 0; i < length; i++ {
		eleList := threeSumTarget(nums[i+1:], target-nums[i])

		for _, ele := range eleList {
			ele = append(ele, nums[i])
			res = append(res, ele)
		}
		for i < length-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
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
// 快慢指针
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
// 快慢指针
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

// 105. 从前序与中序遍历序列构造二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	// 节点为0
	if len(preorder) == 0 {
		return nil
	}
	// 前序遍历第一个元素为根节点
	root := &TreeNode{preorder[0], nil, nil}

	// 在中序遍历中查找根节点位置
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}

	// 递归构建左右子树
	// 参数为左子数的前序遍历，中序遍历
	// 前序遍历的左子树需要通过中序遍历计算出的长度确定
	index := len(inorder[:i]) + 1
	root.Left = buildTree(preorder[1:index], inorder[:i])

	// 参数为右子数的前序遍历，中序遍历
	root.Right = buildTree(preorder[index:], inorder[i+1:])
	return root
}

// 124. 二叉树中的最大路径和
func maxPathSum(root *TreeNode) int {
	minNum := -1 << 20
	// math.MinInt32
	var oneSizeMax func(root *TreeNode) int
	oneSizeMax = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := max(0, oneSizeMax(root.Left))
		right := max(0, oneSizeMax(root.Right))

		minNum = max(minNum, left+right+root.Val)

		log.Print(left, right, root.Val)

		return max(left, right) + root.Val
	}
	oneSizeMax(root)
	return minNum
}

// 230. 二叉搜索树中第K小的元素
func kthSmallest(root *TreeNode, k int) int {
	// 二叉搜索树 中序遍历是有序的
	// 左 中 右
	index, num := 0, 0

	var scanNode func(root *TreeNode)
	scanNode = func(root *TreeNode) {
		if root == nil {
			return
		}
		scanNode(root.Left)
		index++
		if k == index {
			num = root.Val
			return
		}
		scanNode(root.Right)
	}
	scanNode(root)
	return num
}

// 76. 最小覆盖子串
func minWindowV(s string, t string) string {
	need := make(map[byte]int, 0)
	for _, v := range []byte(t) {
		need[v]++
	}
	// 初始化窗口
	window := make(map[byte]int, 0)

	// 有效数量
	number, maxNum := 0, 1<<31-1
	left, right := 0, 0
	start, length := 0, maxNum
	for right < len(s) {
		// 增加窗口
		cur := s[right]
		if _, ok := need[cur]; ok {
			window[cur]++
			// 有效值相等
			if window[cur] == need[cur] {
				number++
			}
		}

		// 增大窗口
		right++

		// 开始缩进窗口
		for number == len(need) {
			if (right - left) < length {
				start = left
				length = right - left
			}
			delEle := s[left]
			left++

			if _, ok := need[delEle]; ok {
				if window[delEle] == need[delEle] {
					number--
				}
				window[delEle]--
			}
		}
	}
	if length == maxNum {
		return ""
	}
	return s[start : start+length]
}

// 567. 字符串的排列
func checkInclusion(t string, s string) bool {
	need := make(map[byte]int)
	window := make(map[byte]int)
	for _, c := range []byte(t) {
		need[c]++
	}
	left := 0
	right := 0
	valid := 0
	for right < len(s) {
		c := s[right]
		right++
		// 进行窗口内数据的一系列更新
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		log.Println(right, left)
		// 判断左侧窗口是否要收缩
		for right-left >= len(t) {
			// 在这里判断是否找到了合法的子串
			if valid == len(need) {
				return true
			}
			d := s[left]
			left++
			// 进行窗口内数据的一系列更新
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	// 未找到符合条件的子串
	return false
}

func lengthOfLongestSubstringV(s string) int {

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	left, right := 0, 0
	window := make(map[byte]int, 0)

	length := -1 << 10
	for right < len(s) {
		ele := s[right]
		right++
		window[ele]++

		for window[ele] > 1 {
			del_ele := s[left]
			left++
			window[del_ele]--
		}
		length = max(length, right-left)
	}

	return length
}

func mergeTwoListsV(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	p1 := l1
	p2 := l2

	for p1 != nil && p2 != nil {
		if p1.Val > p2.Val {
			temp := p2.Next
			res.Next = p2
			p2 = temp
		} else {
			temp := p1.Next
			res.Next = p1
			p1 = temp
		}
	}

	return nil
}
