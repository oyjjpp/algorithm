package hot100

import (
	"container/heap"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"sort"
	"strconv"
	"strings"
	"unicode"
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

// 106. 从中序与后序遍历序列构造二叉树
func buildTreeX(inorder []int, postorder []int) *TreeNode {
	idxMap := map[int]int{}
	for i, v := range inorder {
		idxMap[v] = i
	}
	var build func(int, int) *TreeNode
	build = func(inorderLeft, inorderRight int) *TreeNode {
		// 无剩余节点
		if inorderLeft > inorderRight {
			return nil
		}

		// 后序遍历的末尾元素即为当前子树的根节点
		val := postorder[len(postorder)-1]
		postorder = postorder[:len(postorder)-1]
		root := &TreeNode{Val: val}

		// 根据 val 在中序遍历的位置，将中序遍历划分成左右两颗子树
		// 由于我们每次都从后序遍历的末尾取元素，所以要先遍历右子树再遍历左子树
		inorderRootIndex := idxMap[val]
		root.Right = build(inorderRootIndex+1, inorderRight)
		root.Left = build(inorderLeft, inorderRootIndex-1)
		return root
	}
	return build(0, len(inorder)-1)
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

// 206. 反转链表
// 递归
func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := reverse(head.Next)
	// 把当前节点的子节点的子节点指向当前节点
	// 1>2>3>4<>5
	head.Next.Next = head
	// 头结点变成了尾节点，所以需要置空
	// 1>2>3>4<5
	head.Next = nil

	return last
}

// 206. 反转链表
// 迭代
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	p := head
	for p != nil {
		next := p.Next
		p.Next = pre
		pre = p
		p = next
	}
	return pre
}

// 234. 回文链表
func isPalindromeList(head *ListNode) bool {
	right := head
	var traverse func(*ListNode) bool
	traverse = func(node *ListNode) bool {
		if node == nil {
			return true
		}
		res := traverse(node.Next)
		// 后序遍历代码
		res = res && (node.Val == right.Val)
		right = right.Next
		return res
	}

	return traverse(head)
}

// 反转链表的前N个元素
// 保存后续
var nextList *ListNode

func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		nextList = head.Next
		return head
	}

	last := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = nextList
	return last
}

// 92. 反转链表 II
// 1,2,3,4,5,6
// 3,5
// 1,2,5,4,3,6
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return reverseN(head, right)
	}

	// 前进到反转的起点触发 base case
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

// 25. K 个一组翻转链表
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	a, b := head, head

	// 迭代出k个 元素
	for i := 0; i < k; i++ {
		// 如果长度不足直接返回
		if b == nil {
			return head
		}
		b = b.Next
	}
	// 先反转以head开头的k的元素
	newHead := reverseRange(a, b)
	// 将第 k + 1 个元素作为 head 递归调用 reverseKGroup 函数。
	a.Next = reverseKGroup(b, k)
	return newHead
}

// 反转一定范围的链表
func reverseRange(a, b *ListNode) *ListNode {
	var pre *ListNode
	cur := a
	// while 终止的条件改一下就行了
	for cur != b {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	// 返回反转后的头结点
	return pre
}

// 111. 二叉树的最小深度
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	minPath := 1

	for len(queue) > 0 {
		sz := len(queue)

		for i := 0; i < sz; i++ {
			node := queue[i]
			log.Println(i, node.Val)
			// 碰到叶子节点
			if node.Left == nil && node.Right == nil {
				return minPath
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[sz:]
		minPath++
	}
	return minPath
}

// 752. 打开转盘锁
func openLock(deadends []string, target string) int {
	var plusOne func(s string, j int) string
	plusOne = func(s string, j int) string {
		ch := []byte(s)
		if ch[j] == '9' {
			ch[j] = '0'
		} else {
			ch[j] += 1
		}
		return string(ch)
	}

	var minusOne func(s string, j int) string
	minusOne = func(s string, j int) string {
		ch := []byte(s)
		if ch[j] == '0' {
			ch[j] = '9'
		} else {
			ch[j] -= 1
		}
		return string(ch)
	}
	// 死亡密码
	deads := make(map[string]bool)
	for _, s := range deadends {
		deads[s] = true
	}

	// 记录已经穷举过的密码，防止走回头路
	visited := make(map[string]bool)

	queue := make([]string, 0)
	queue = append(queue, "0000")
	visited["0000"] = true

	step := 0

	for len(queue) > 0 {
		sz := len(queue)

		for i := 0; i < sz; i++ {
			node := queue[0]
			queue = queue[1:]

			/* 判断是否到达终点 */
			if _, ok := deads[node]; ok {
				continue
			}
			if node == target {
				return step
			}

			for j := 0; j < 4; j++ {
				up := plusOne(node, j)
				if _, ok := visited[up]; !ok {
					queue = append(queue, up)
					visited[up] = true
				}
				down := minusOne(node, j)
				if _, ok := visited[down]; !ok {
					queue = append(queue, down)
					visited[down] = true
				}
			}
		}
		step++
	}
	return -1
}

// 向上拨动
func plusOne(s string, j int) string {
	ch := []byte(s)
	if ch[j] == '9' {
		ch[j] = '0'
	} else {
		ch[j] += 1
	}
	return string(ch)
}

// 向下拨动
// 将 s[i] 向下拨动一次
func minusOne(s string, j int) string {
	ch := []byte(s)
	if ch[j] == '0' {
		ch[j] = '9'
	} else {
		ch[j] -= 1
	}
	return string(ch)
}

func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + right

		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			right = mid - 1
		}
	}
	log.Println(left)
	if left == len(nums) {
		return []int{-1, -1}
	}

	if nums[left] == target {
		for i := left; i < len(nums); i++ {
			if nums[i] > target {
				return []int{left, i - 1}
			}
		}
		return []int{left, len(nums) - 1}
	}
	return []int{-1, -1}
}

// 33. 搜索旋转排序数组
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

// 46. 全排列
func permute(nums []int) [][]int {
	// 存储结果集
	res := make([][]int, 0)
	// 组合元素
	track := make([]int, 0)
	used := make([]bool, len(nums))

	var backtrack func(track []int, used []bool)
	backtrack = func(track []int, used []bool) {
		if len(track) == len(nums) {
			temp := make([]int, len(nums))
			copy(temp, track)
			res = append(res, temp)
			return
		}

		for i := range nums {
			// 排除不合法的选择
			if used[i] {
				continue
			}

			// 做选择
			track = append(track, nums[i])
			used[i] = true
			backtrack(track, used)

			// 撤销选择
			track = track[:len(track)-1]
			used[i] = false
		}
	}
	backtrack(track, used)
	return res
}

// 47. 全排列 II
func permuteUnique(nums []int) [][]int {
	// 存储结果集
	res := make([][]int, 0)
	// 组合元素
	track := make([]int, 0)
	used := make([]bool, len(nums))

	var backtrack func(track []int, used []bool)
	backtrack = func(track []int, used []bool) {
		if len(track) == len(nums) {
			temp := make([]int, len(nums))
			copy(temp, track)
			res = append(res, temp)
			return
		}

		for i := range nums {
			// 排除不合法的选择
			if used[i] {
				continue
			}

			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}

			// 做选择
			track = append(track, nums[i])
			used[i] = true
			backtrack(track, used)

			// 撤销选择
			track = track[:len(track)-1]
			used[i] = false
		}
	}
	sort.Ints(nums)
	backtrack(track, used)
	return res
}

// 排列（元素无重可复选）
func permuteRepeat(nums []int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)
	var backtrack func()
	backtrack = func() {
		if len(track) == len(nums) {
			temp := make([]int, len(nums))
			copy(temp, track)
			res = append(res, temp)
			return
		}

		for i := 0; i < len(nums); i++ {
			track = append(track, nums[i])
			backtrack()
			track = track[:len(track)-1]
		}
	}
	backtrack()
	return res
}

// 51. N 皇后 校验函数
func solveNQueens(n int) [][]string {
	// 校验是否为有效范围
	var isValid func(board []string, row, col int) bool
	isValid = func(board []string, row, col int) bool {
		// 检查列
		for i := 0; i < row; i++ {
			if board[i][col] == 'Q' {
				return false
			}
		}
		n := len(board)
		// 右上方
		for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
			if board[i][j] == 'Q' {
				return false
			}
		}
		for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
			if board[i][j] == 'Q' {
				return false
			}
		}
		return true
	}

	// 保存结果
	res := make([][]string, 0)
	board := make([]string, n)
	for i := 0; i < n; i++ {
		board[i] = strings.Repeat(".", n)
	}

	var backtrack func(board []string, row int)
	backtrack = func(board []string, row int) {
		if row == len(board) {
			newRow := make([]string, len(board))
			copy(newRow, board)
			res = append(res, newRow)
			return
		}

		n := len(board[row])
		for col := 0; col < n; col++ {
			// 校验是否符合要求
			if !isValid(board, row, col) {
				continue
			}

			newLine := []byte(board[row])
			newLine[col] = 'Q'
			board[row] = string(newLine)

			backtrack(board, row+1)

			newLine[col] = '.'
			board[row] = string(newLine)
		}
	}

	backtrack(board, 0)
	return res
}

// 78. 子集
func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)
	var backtrack func(start int)
	backtrack = func(start int) {
		temp := make([]int, len(track))
		copy(temp, track)
		res = append(res, temp)

		// 使用start 保证子集
		for i := start; i < len(nums); i++ {
			track = append(track, nums[i])
			log.Println(i, track)
			backtrack(i + 1)
			track = track[:len(track)-1]
		}
	}
	backtrack(0)
	return res
}

// 子集 II
func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)

	// 先排序，让相同的元素靠在一起
	sort.Ints(nums)

	var backtrack func(start int)
	backtrack = func(start int) {
		temp := make([]int, len(track))
		copy(temp, track)
		res = append(res, temp)

		// 使用start 保证子集
		for i := start; i < len(nums); i++ {

			// 值相同的 相邻节点 直接过滤
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			track = append(track, nums[i])
			log.Println(i, track)
			backtrack(i + 1)
			track = track[:len(track)-1]
		}
	}
	backtrack(0)
	return res
}

// 77. 组合
func combine(n int, k int) [][]int {
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, i+1)
	}
	res := make([][]int, 0)
	track := make([]int, 0)
	var backtrack func(start int)
	backtrack = func(start int) {
		if len(track) == k {
			temp := make([]int, k)
			copy(temp, track)
			res = append(res, temp)
			return
		}

		// 使用start 保证子集
		for i := start; i < n; i++ {
			track = append(track, nums[i])
			backtrack(i + 1)
			track = track[:len(track)-1]
		}
	}
	backtrack(0)
	return res
}

// 39. 组合总和
// 子集/组合（元素无重可复选
func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)
	trackSum := 0

	var backtrack func(start int)
	backtrack = func(start int) {
		if trackSum == target {
			temp := make([]int, len(track))
			copy(temp, track)
			res = append(res, temp)
		}
		if trackSum > target {
			return
		}

		for i := start; i < len(candidates); i++ {
			trackSum += candidates[i]
			track = append(track, candidates[i])

			backtrack(i)

			trackSum -= candidates[i]
			track = track[:len(track)-1]
		}
	}
	backtrack(0)
	return res
}

func maxDepth(root *TreeNode) int {
	var dp func(root *TreeNode) int
	dp = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := dp(root.Left)
		right := dp(root.Right)

		return max(left, right) + 1
	}
	return dp(root)
}

// 扫描二叉树节点所在层次
func printBinaryLevel(root *TreeNode) {
	if root == nil {
		return
	}

	var dp func(root *TreeNode, level int)
	dp = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		dp(root.Left, level+1)
		log.Printf("node:%d,level:%d", root.Val, level)
		dp(root.Right, level+1)

	}
	dp(root, 0)
}

// 543. 二叉树的直径
func diameterOfBinaryTree(root *TreeNode) int {
	var dp func(root *TreeNode) int

	// 可能是左右加起来最大 maxNumber
	maxNumber := 0
	dp = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := dp(root.Left)
		right := dp(root.Right)
		maxNumber = max(maxNumber, left+right)
		return max(left, right) + 1
	}

	dp(root)
	return maxNumber
}

// 515. 在每个树行中找最大值
func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	data := make([]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		sz := len(queue)

		maxNumber := -1 << 10
		for i := 0; i < sz; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Val > maxNumber {
				maxNumber = node.Val
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		data = append(data, maxNumber)
	}
	return data
}

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

// 26. 删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}

// 83. 删除排序链表中的重复元素
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	for fast != nil {
		if slow.Val != fast.Val {
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}
	slow.Next = nil
	return head
}

// 27. 移除元素
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

// 283. 移动零
func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	// 寻找所有等于0的数据，进行移除
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}

	// 替换后面非零的数据
	for ; slow < len(nums); slow++ {
		nums[slow] = 0
	}
}

// 167. 两数之和 II - 输入有序数组
func twoSumTarget(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else if sum > target {
			right--
		}
	}
	return []int{-1, -1}
}

// 344. 反转字符串
func reverseString(s []byte) {
	left, right := 0, len(s)-1

	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// 125. 验证回文串
// 双指针
func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	// 判断是否合法的字符
	isValid := func(v rune) bool {
		return unicode.IsDigit(v) || unicode.IsLetter(v)
	}

	s = strings.ToLower(s)
	str := []rune(s)
	slow, fast := 0, len(s)-1
	for slow < fast {
		// 不是字符串
		if !isValid(str[slow]) {
			slow++
			continue
		}

		// 验证是否字符串
		if !isValid(str[fast]) {
			fast--
			continue
		}

		if str[slow] != str[fast] {
			return false
		}
		slow++
		fast--
	}
	return true
}

// 5. 最长回文子串
// 从中心向两端扩散的双指针技巧
func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		s1 := palindrome(s, i, i)
		s2 := palindrome(s, i, i+1)

		if len(res) < len(s1) {
			res = s1
		}
		if len(res) < len(s2) {
			res = s2
		}
	}
	return res
}

func palindrome(s string, l int, r int) string {
	// 防止索引越界
	for l >= 0 && r < len(s) && s[l] == s[r] {
		// 向两边展开
		l--
		r++
	}
	// 返回以 s[l] 和 s[r] 为中心的最长回文串
	return s[l+1 : r]
}

// 19. 删除链表的倒数第 N 个结点
// 通过双指针找到该节点的位置
// 然后删除节点
func removeNthFromEndV(head *ListNode, n int) *ListNode {
	newNode := &ListNode{0, head}
	left, right := newNode, head

	index := 0
	for right != nil {
		index++
		if index > n {
			left = left.Next
		}
		right = right.Next
	}

	left.Next = left.Next.Next
	return newNode.Next
}

func getIntersectionNodeV(headA, headB *ListNode) *ListNode {
	left, right := headA, headB
	for left != right {
		if left == nil {
			left = headB
		} else {
			left = left.Next
		}

		if right == nil {
			right = headA
		} else {
			right = right.Next
		}
	}
	return left
}

// 动态规划
// 509. 斐波那契数
func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

// 自顶向下【备忘录】
// 509. 斐波那契数
func fibV(n int) int {
	origin := map[int]int{}

	var dp func(data map[int]int, m int) int

	dp = func(data map[int]int, m int) int {
		if m == 0 || m == 1 {
			return m
		}

		if data[m] != 0 {
			return data[m]
		}
		data[m] = dp(data, m-1) + dp(data, m-2)
		return data[m]
	}
	return dp(origin, n)
}

// 自底向上
// 509. 斐波那契数
func fibV2(n int) int {
	if n == 0 {
		return 0
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 动态规划
// 322. 零钱兑换
// 自顶向下
func coinChange(coins []int, amount int) int {
	memo := make([]int, amount+1)
	// 将备忘录初始化为 -666，代表还未被计算
	for i := range memo {
		memo[i] = -666
	}
	maxNum := 1 << 20
	fmt.Println(maxNum)

	var dp func(coins []int, amount int) int
	dp = func(coins []int, amount int) int {
		if amount == 0 {
			return 0
		}
		if amount < 0 {
			return -1
		}

		if memo[amount] != -666 {
			return memo[amount]
		}

		res := maxNum
		for _, coin := range coins {
			subProblem := dp(coins, amount-coin)
			if subProblem == -1 {
				continue
			}
			res = min(res, 1+subProblem)
		}
		if res == maxNum {
			memo[amount] = -1
		} else {
			memo[amount] = res
		}
		return memo[amount]
	}

	return dp(coins, amount)
}

// 动态规划
// 322. 零钱兑换
// 自底向上
func coinChangeV(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = amount + 1
	}

	dp[0] = 0
	for i := 0; i < len(dp); i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

// 状态 、 选择
// 状态 -> 选择 （穷举）

func maxProfit(prices []int) int {
	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	k := 2
	days := len(prices)
	data := make([][][]int, days)
	for i := range data {
		data[i] = make([][]int, k+1)
		for j := range data[i] {
			data[i][j] = make([]int, 2)
		}
	}

	for i := 0; i < days; i++ {
		if i == 0 {
			data[i][k][0] = 0
			data[i][k][1] = -prices[i]
			continue
		}

		data[i][k][0] = max(data[i-1][k][0], data[i-1][k][1]+prices[i])
		data[i][k][1] = max(data[i-1][k][1], data[i-1][k-1][0]-prices[i])
	}
	return data[days-1][k][0]
}

// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码还未经过力扣测试，仅供参考，如有疑惑，可以参照我写的 java 代码对比查看。

// 原始版本
func maxProfit_k_2(prices []int) int {

	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	max_k := 2 // 最大可交易次数
	n := len(prices)
	dp := make([][][]int, n) // i为天数，k为当前第几次交易，0表示不持有股票，1表示持有股票
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, max_k+1)
		for k := 0; k < max_k+1; k++ {
			dp[i][k] = make([]int, 2)
		}
	}
	for i := 0; i < n; i++ {
		for k := max_k; k >= 1; k-- {
			if i-1 == -1 {
				// 处理 base case
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[i]
				continue
			}
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}
	// 穷举了 n × max_k × 2 个状态，正确。
	return dp[n-1][max_k][0] //返回最大利润
}

func maxProfit_v(prices []int) int {
	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	k := 2
	n := len(prices)
	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, k+1) // 构建 3 维 DP 数组
		for j := range dp[i] {
			dp[i][j] = make([]int, 2) // 初始化 DP 数组
		}
	}

	for i := 0; i < n; i++ {
		for j := k; j >= 1; j-- {
			if i-1 == -1 {
				// 处理 base case
				dp[i][j][0] = 0
				dp[i][j][1] = -prices[i]
				continue
			}
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
	}

	return dp[n-1][k][0]
}

// 状态
// 选择 偷、不偷

// 198. 打家劫舍
func rob(nums []int) int {
	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	data := map[int]int{}
	res := 0
	var dp func(number []int, start int) int
	dp = func(number []int, start int) int {
		if start >= len(number) {
			return 0
		}

		if temp, ok := data[start]; ok {
			return temp
		}
		// 不去抢
		res = max(dp(number, start+1), dp(number, start+2)+number[start])
		data[start] = res
		return res
	}
	res = dp(nums, 0)
	return res
}

// 198. 打家劫舍
func rob_v1(nums []int) int {
	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	length := len(nums)
	data := make([]int, length+2)

	for i := length - 1; i >= 0; i-- {
		data[i] = max(data[i+1], data[i+2]+nums[i])
	}
	return data[0]
}

// 213. 打家劫舍 II
func rob_v2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var dp func(number []int) int
	dp = func(number []int) int {
		length := len(number)
		data := make([]int, length+2)

		for i := length - 1; i >= 0; i-- {
			data[i] = max(data[i+1], data[i+2]+number[i])
		}
		return data[0]
	}
	temp1 := nums[0 : len(nums)-1]
	temp2 := nums[1:]
	fmt.Println(temp1)
	fmt.Println(temp2)
	res := max(dp(temp1), dp(temp2))
	return res
}

// 337. 打家劫舍 III
func rob_v3(root *TreeNode) int {
	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	data := make(map[*TreeNode]int, 0)
	var dp func(node *TreeNode) int
	dp = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		if temp, ok := data[node]; ok {
			return temp
		}

		// 不抢
		notDo := dp(node.Left) + dp(node.Right)

		// 抢
		doIt := node.Val
		if node.Left != nil {
			doIt += dp(node.Left.Left) + dp(node.Left.Right)
		}

		if node.Right != nil {
			doIt += dp(node.Right.Left) + dp(node.Right.Right)
		}

		res := max(notDo, doIt)
		data[node] = res
		return res
	}

	res := dp(root)
	return res
}

// 303. 区域和检索 - 数组不可变
type NumArray struct {
	Number []int
}

func Constructor(nums []int) NumArray {
	temp := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		temp[i] = temp[i-1] + nums[i-1]
	}

	return NumArray{temp}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.Number[right+1] - this.Number[left]
}

// 304. 二维区域和检索 - 矩阵不可变
type NumMatrix struct {
	Number [][]int
}

func ConstructorA(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	if m == 0 || n == 0 {
		return NumMatrix{}
	}
	number := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		number[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			number[i][j] = number[i-1][j] + number[i][j-1] + matrix[i-1][j-1] - number[i-1][j-1]
		}
	}
	return NumMatrix{number}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.Number[row2+1][col2+1] - this.Number[row1][col2+1] - this.Number[row2+1][col1] + this.Number[row1][col1]
}

// superEggDrop
// 887. 鸡蛋掉落
func superEggDrop(K int, N int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	// 备忘录解决重叠子问题
	meno := map[string]int{}

	var dp func(K, N int) int
	dp = func(K, N int) int {
		// base case
		// 楼层为0
		if N == 0 {
			return 0
		}
		// 只有一个鸡蛋
		if K == 1 {
			return N
		}

		// 校验备忘录中是否存在
		key := fmt.Sprintf("%d%d", K, N)
		if rs, ok := meno[key]; ok {
			return rs
		}

		res := 1<<31 - 1
		left, right := 1, N
		for left <= right {
			mid := (left + right) / 2

			// 状态
			// 碎了
			broken := dp(K-1, mid-1)
			// 没碎
			notBroken := dp(K, N-mid)

			// 选择

			if broken > notBroken {
				right = mid - 1
				res = min(res, broken+1)
			} else {
				left = mid + 1
				res = min(res, notBroken+1)
			}
		}
		meno[key] = res
		return res
	}
	return dp(K, N)
}

func TestFunc() {
	// 0,1,2,3,4
	// 8,5,9,6,1
	// 1,2,3,4,5
	nums := []int{8, 5, 9, 6, 1}

	diff := make([]int, len(nums))
	// 构造差分数组
	diff[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		diff[i] = nums[i] - nums[i-1]
	}

	res := make([]int, len(diff))
	// 根据差分数组构造结果数组
	res[0] = diff[0]
	for i := 1; i < len(diff); i++ {
		res[i] = res[i-1] + diff[i]
	}
}

// 差分分组
type Difference struct {
	diff []int
}

// 输入一个数组，初始化其差分分组
func NewDifference(nums []int) *Difference {
	diff := make([]int, len(nums))
	diff[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		diff[i] = nums[i] - nums[i-1]
	}
	return &Difference{diff: diff}
}

// 返回原数组
func (d *Difference) Result() []int {
	res := make([]int, len(d.diff))
	res[0] = d.diff[0]
	for i := 1; i < len(d.diff); i++ {
		res[i] = d.diff[i] + res[i-1]
	}
	return res
}

// 给闭区间 [i, j] 增加 val（可以是负数）
func (d *Difference) Increment(i, j, val int) {
	d.diff[i] += val
	if j+1 < len(d.diff) {
		d.diff[j+1] -= val
	}
}

// 370 题「 区间加法」 就
func getModifiedArray(length int, updates [][]int) []int {
	// nums 初始化为全 0
	nums := make([]int, length)
	// 构造差分解法
	df := NewDifference(nums)

	for _, update := range updates {
		i, j, val := update[0], update[1], update[2]
		df.Increment(i, j, val)
	}

	return df.Result()
}

// 1109. 航班预订统计
func corpFlightBookings(bookings [][]int, n int) []int {
	// nums 初始化为全 0
	nums := make([]int, n)
	// 构造差分解法
	df := NewDifference(nums)

	for _, booking := range bookings {
		// 注意转成数组索引要减一哦
		i := booking[0] - 1
		j := booking[1] - 1
		val := booking[2]
		// 对区间 nums[i..j] 增加 val
		df.Increment(i, j, val)
	}
	// 返回最终的结果数组
	return df.Result()
}

// 1094. 拼车
func carPooling(trips [][]int, capacity int) bool {
	// 最多有 1001 个车站
	nums := make([]int, 1001)
	// 构造差分解法
	df := NewDifference(nums)

	for _, trip := range trips {
		// 乘客数量
		val := trip[0]
		// 第 trip[1] 站乘客上车
		i := trip[1]
		// 第 trip[2] 站乘客已经下车，
		// 即乘客在车上的区间是 [trip[1], trip[2] - 1]
		j := trip[2] - 1
		// 进行区间操作
		df.Increment(i, j, val)
	}

	res := df.Result()

	// 客车自始至终都不应该超载
	for i := 0; i < len(res); i++ {
		if capacity < res[i] {
			return false
		}
	}
	return true
}

// 48. 旋转图像
func rotate(matrix [][]int) {
	length := len(matrix)

	// 先沿着对角线镜像对称二维数组
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	reverse := func(data []int) {
		left, right := 0, len(data)-1

		for left < right {
			data[left], data[right] = data[right], data[left]
			left++
			right--
		}
	}

	for _, node := range matrix {
		reverse(node)
	}
}

// 54. 螺旋矩阵
func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	upper_bound, lower_bound := 0, m-1

	left_bound, right_bound := 0, n-1
	res := make([]int, 0, m*n)
	// len(res) == m * n 则遍历完整个数组
	for len(res) < m*n {
		if upper_bound <= lower_bound {
			// 在顶部从左向右遍历
			for j := left_bound; j <= right_bound; j++ {
				res = append(res, matrix[upper_bound][j])
			}
			// 上边界下移
			upper_bound++
		}

		if left_bound <= right_bound {
			// 在右侧从上向下遍历
			for i := upper_bound; i <= lower_bound; i++ {
				res = append(res, matrix[i][right_bound])
			}
			// 右边界左移
			right_bound--
		}

		if upper_bound <= lower_bound {
			// 在底部从右向左遍历
			for j := right_bound; j >= left_bound; j-- {
				res = append(res, matrix[lower_bound][j])
			}
			// 下边界上移
			lower_bound--
		}

		if left_bound <= right_bound {
			// 在左侧从下向上遍历
			for i := lower_bound; i >= upper_bound; i-- {
				res = append(res, matrix[i][left_bound])
			}
			// 左边界右移
			left_bound++
		}
	}
	return res
}

// 59. 螺旋矩阵 II
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	upper_bound, lower_bound := 0, n-1
	left_bound, right_bound := 0, n-1
	// 需要填入矩阵的数字
	num := 1

	for num <= n*n {
		if upper_bound <= lower_bound {
			// 在顶部从左向右遍历
			for j := left_bound; j <= right_bound; j++ {
				matrix[upper_bound][j] = num
				num++
			}
			// 上边界下移
			upper_bound++
		}

		if left_bound <= right_bound {
			// 在右侧从上向下遍历
			for i := upper_bound; i <= lower_bound; i++ {
				matrix[i][right_bound] = num
				num++
			}
			// 右边界左移
			right_bound--
		}

		if upper_bound <= lower_bound {
			// 在底部从右向左遍历
			for j := right_bound; j >= left_bound; j-- {
				matrix[lower_bound][j] = num
				num++
			}
			// 下边界上移
			lower_bound--
		}

		if left_bound <= right_bound {
			// 在左侧从下向上遍历
			for i := lower_bound; i >= upper_bound; i-- {
				matrix[i][left_bound] = num
				num++
			}
			// 左边界右移
			left_bound++
		}
	}
	return matrix
}

// 187. 重复的DNA序列
func findRepeatedDnaSequences(s string) []string {
	length := len(s)
	res := make(map[string]bool, 0)
	scan := make(map[string]bool, 0)

	for i := 0; i+10 <= length; i++ {
		subStr := s[i : i+10]

		if _, ok := scan[subStr]; ok {
			res[subStr] = true
		} else {
			scan[subStr] = true
		}
	}

	data := make([]string, 0)
	for key := range res {
		data = append(data, key)
	}
	return data
}

// 剑指 Offer 53 - I. 在排序数组中查找数字 I
func searchOffer53(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		log.Println("for", left, mid, right)
		if target == nums[mid] {
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		}
	}

	log.Println(left)
	if left == len(nums) {
		return 0
	}

	if nums[left] != target {
		return 0
	}

	number := 0

	for ; left < len(nums); left++ {
		if nums[left] == target {
			number++
		} else {
			break
		}
	}
	return number
}

// 528. 按权重随机选择
type Solution struct {
	preSum []int
}

func ConstructorSolution(w []int) *Solution {
	preSum := make([]int, len(w)+1)

	preSum[0] = 0
	for i := 1; i <= len(w); i++ {
		preSum[i] = w[i-1] + preSum[i-1]
	}
	log.Println(preSum)
	return &Solution{preSum}
}

func (this *Solution) PickIndex() int {
	n := len(this.preSum)

	result, _ := rand.Int(rand.Reader, big.NewInt(int64(this.preSum[n-1])))
	index := result.Int64() + 1
	log.Println(n, this.preSum, index)
	return this.sesrch(int(index)) - 1
}

func (this *Solution) sesrch(target int) int {
	if len(this.preSum) == 0 {
		return -1
	}
	// [0,1,2]
	// [0,1,4]
	left, right := 0, len(this.preSum)
	for left < right {
		mid := left + (right-left)/2
		if this.preSum[mid] == target {
			right = mid
		} else if this.preSum[mid] < target {
			left = mid + 1
		} else if this.preSum[mid] > target {
			right = mid
		}
	}
	return left
}

// 870. 优势洗牌
func advantageCount(nums1 []int, nums2 []int) []int {
	n := len(nums1)
	idx1 := make([]int, n)
	idx2 := make([]int, n)
	for i := 1; i < n; i++ {
		idx1[i] = i
		idx2[i] = i
	}

	// 根据元素对索引位置进行大小排序
	sort.Slice(idx1, func(i, j int) bool {
		return nums1[idx1[i]] < nums1[idx1[j]]
	})
	sort.Slice(idx2, func(i, j int) bool {
		return nums2[idx2[i]] < nums2[idx2[j]]
	})

	log.Println(idx1, idx2)
	log.Println(nums1, nums2)

	// 保存结果
	ans := make([]int, n)
	left, right := 0, n-1

	for i := 0; i < n; i++ {
		if nums1[idx1[i]] > nums2[idx2[left]] {
			ans[idx2[left]] = nums1[idx1[i]]
			left++
		} else {
			ans[idx2[right]] = nums1[idx1[i]]
			right--
		}
	}
	return ans
}

// 870. 优势洗牌
func advantageCountX(nums1 []int, nums2 []int) []int {
	return nil
}

type PriorityArray [][]int

func (pq PriorityArray) Len() int {
	return len(pq)
}

func (pq PriorityArray) Less(i, j int) bool {
	return pq[i][1] > pq[j][1]
}

func (pq PriorityArray) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityArray) Push(x interface{}) {
	*pq = append(*pq, x.([]int))
}

func (pq *PriorityArray) Pop() interface{} {
	n := len(*pq)
	item := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return item
}

// bubblingSort
// 冒泡排序
func bubblingSort(data []int) {
	length := len(data)
	if length < 2 {
		return
	}

	// 外层循环，告知排序得次数
	for i := 0; i < length-1; i++ {
		// 内层循环，用户循环比较，每一步骤都将最大得移动到最右侧
		for j := 0; j < length-i-1; j++ {
			log.Println(j)
			// 是否会溢出？
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}

// selectSort
// 选择排序
func selectSort(data []int) {
	length := len(data)
	if length < 2 {
		return
	}
	for i := 0; i < length-1; i++ {
		// 假定最大值索引位置为第一个
		maxpos := 0

		// 内部循环找到真实得最大值索引位置
		for j := 0; j < length-i; j++ {
			if data[j] > data[maxpos] {
				maxpos = j
			}
		}
		data[length-i-1], data[maxpos] = data[maxpos], data[length-i-1]
	}
}

// insertSort
// 插入排序
func insertSort(data []int) {
	length := len(data)
	if length < 2 {
		return
	}

	// 假设第一个元素为有序数组
	for i := 1; i < length; i++ {
		// 选取一个值 向有序集合中插入
		temp := data[i]

		// 将选取的值插入到有序集合的指定位置
		j := i - 1
		// 挑选的值与倒序与有序集合对比，如果选择的值小于则移动有序集合索引
		for ; j >= 0 && data[j] > temp; j-- {
			data[j+1] = data[j]
		}
		data[j+1] = temp
	}
}

/**
 * @desc 归并排序
 * 思路： 选择中间索引将数组分割为两个，然后组合两个数组按大小顺序组合
 * 时间复杂度 O(nlog2n)
 * 空间复杂度 O(n) + O(log2n)
 * 稳定性：稳定
 */
func MergeSort(item []int) []int {
	mergeSort(item, 0, len(item)-1)
	return item
}

// @param item 排序数组
// @param 开始索引位置
// @param 结束索引位置
func mergeSort(item []int, left, right int) {
	if left < right {
		center := (left + right) / 2
		mergeSort(item, left, center)
		mergeSort(item, center+1, right)
		merge(item, left, center+1, right)
	}
}

// @desc 合并两个数组
func merge(item []int, left, center, right int) {
	// 左侧数组大小
	leftData := make([]int, center-left)
	// 右侧数组大小
	rightData := make([]int, right-center+1)

	// 向两个数组中填充数据
	for i := left; i < center; i++ {
		leftData[i-left] = item[i]
	}

	for i := center; i <= right; i++ {
		rightData[i-center] = item[i]
	}

	// 用于遍历两个数组
	i, j := 0, 0
	// 数组中的第一个元素
	index := left
	// 循环对比合并两个数组
	for i < len(leftData) && j < len(rightData) {
		if leftData[i] < rightData[j] {
			item[index] = leftData[i]
			i++
		} else {
			item[index] = rightData[j]
			j++
		}
		// 增加后索引增加1
		index++
	}

	// 将数据中剩余的元素继续插入
	for i < len(leftData) {
		item[index] = leftData[i]
		i++
		index++
	}
	for j < len(rightData) {
		item[index] = rightData[j]
		j++
		index++
	}
}

func mergeSortV2(item []int, left, center, right int) {
	temp := make([]int, right-left+1)
	i := 0
	// 左边开始
	start := left
	// 右边开始
	end := center + 1
	for start <= center && end <= right {
		if item[start] < item[end] {
			temp[i] = item[start]
			start++
		} else {
			temp[i] = item[end]
			end++
		}
		i++
	}

	for start <= center {
		temp[i] = item[start]
		i++
		start++
	}
	for end <= end {
		temp[i] = item[end]
		i++
		end++
	}
	// 将结果返回给原素组
	for i := 0; i < len(temp); i++ {
		item[left+i] = temp[i]
	}
}

// quickSort
// 快速排序，使用递归方式
func quickSort(data []int) []int {
	if len(data) == 0 {
		return data
	}
	// 取第一个元素作为比较节点
	// temp := data[0]
	left := []int{}
	right := []int{}

	// 此处循环从1开始，第一个节点已经拿出来用于比较了
	for i := 1; i < len(data); i++ {
		if data[i] > data[0] {
			right = append(right, data[i])
		} else {
			left = append(left, data[i])
		}
	}
	left = quickSort(left)
	right = quickSort(right)
	return append(append(left, data[0]), right...)
}

// quickSort
// 快速排序，使用递归方式
func quickSortX(data []int, left, right int) {
	if left >= right {
		return
	}

	start := left
	end := right

	// 选择第一个节点
	value := data[left]
	// 根据选取的节点将数组 分成两部分
	for left < right {
		// 如果右侧节点大于选取节点则向左移动指针
		for right > left && data[right] >= value {
			right--
		}
		data[left] = data[right]

		// 如果左边节点小于选取节点则向右移动指针
		for left < right && data[left] <= value {
			left++
		}
		data[right] = data[left]
	}
	// 将选取得值赋值左指针
	data[left] = value
	quickSortX(data, start, left-1)
	quickSortX(data, left+1, end)
}

// 扫描二叉树节点数
func ScanTreeCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftCount := ScanTreeCount(root.Left)
	rightCount := ScanTreeCount(root.Right)
	// 后序位置
	fmt.Printf("节点 %v 的左子树有 %d 个节点，右子树有 %d 个节点 \n",
		root, leftCount, rightCount)

	return leftCount + rightCount + 1
}

// 226. 翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		temp := root.Left
		root.Left = root.Right
		root.Right = temp

		traverse(root.Left)
		traverse(root.Right)
	}
	traverse(root)
	return root
}

// 116. 填充每个节点的下一个右侧节点指针
func connect(root *Node) *Node {
	var traverse func(left, right *Node)
	traverse = func(left, right *Node) {
		if left == nil || right == nil {
			return
		}

		left.Next = right
		traverse(left.Left, left.Right)
		traverse(right.Left, right.Right)
		traverse(left.Right, right.Left)
	}
	if root == nil {
		return root
	}

	traverse(root.Left, root.Right)
	return root
}

// 114. 二叉树展开为链表
func flatten(root *TreeNode) {
	// base case
	if root == nil {
		return
	}

	// 利用定义，把左右子树拉平
	flatten(root.Left)
	flatten(root.Right)

	/**** 后序遍历位置 ****/
	// 1、左右子树已经被拉平成一条链表
	left := root.Left
	right := root.Right

	// 2、将左子树作为右子树
	root.Left = nil
	root.Right = left

	// 3、将原先的右子树接到当前右子树的末端
	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right
}

// 654. 最大二叉树
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	maxValue, index := -1<<10, 0

	for i := 0; i < len(nums); i++ {
		if maxValue < nums[i] {
			maxValue = nums[i]
			index = i
		}
	}

	root := &TreeNode{
		Val: maxValue,
	}
	root.Left = constructMaximumBinaryTree(nums[:index])
	root.Right = constructMaximumBinaryTree(nums[index+1:])
	return root
}

// 297. 二叉树的序列化与反序列化
type Codec struct{}

func CodecInit() (_ Codec) {
	return
}

func (Codec) serialize(root *TreeNode) string {
	sb := &strings.Builder{}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		// 前序遍历 将数据存储到builder中
		if node == nil {
			sb.WriteString("null,")
			return
		}
		sb.WriteString(strconv.Itoa(node.Val))
		sb.WriteByte(',')

		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return sb.String()
}

func (Codec) deserialize(data string) *TreeNode {
	sp := strings.Split(data, ",")
	var build func() *TreeNode
	build = func() *TreeNode {
		if sp[0] == "null" {
			sp = sp[1:]
			return nil
		}
		val, _ := strconv.Atoi(sp[0])
		sp = sp[1:]
		return &TreeNode{val, build(), build()}
	}
	return build()
}

// 652. 寻找重复的子树
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {}
