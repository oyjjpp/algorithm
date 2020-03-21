package leetcode

// 归并排序方式
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)

	rs := make([]int, 0, length)
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			rs = append(rs, nums1[i])
			i++
		} else {
			rs = append(rs, nums2[j])
			j++
		}
	}

	// num2 还存在剩余
	if i >= len(nums1) {
		rs = append(rs, nums2[j:]...)
	}

	// num1 还存在剩余
	if j >= len(nums2) {
		rs = append(rs, nums1[i:]...)
	}

	if length%2 == 0 {
		return float64(rs[length/2]+rs[length/2-1]) / 2
	}
	return float64(rs[length/2])
}

// findMedianSortedArrays2 给出两个有序数组，假设两个数组的长度和是len，
// 如果len为奇数，那么我们求的就是两个数组合并后的第 (len >> 1) + 1 大的数，
// 如果len为偶数，就是第 (len >> 1) 和 (len >> 1) + 1 两个数的平均数
// 给定两个有序数组，求第k大数,如果我们从 A 和 B 中分别取前 k/2 个元素，其中必然有一部分是是在数组 C 的前 k 个数里
// 设 mid = k / 2，当 A[mid - 1] < B[mid - 1] 时，可以断定 A 的前 mid 个元素是在 C 的前 k 个数里
// 时间复杂度是 O(log(m+n))
func findMedianSortedArraysV2(nums1 []int, nums2 []int) {
	// 假设第一个数组长度较短
	m, n := len(nums1), len(nums2)
	if m > n {
		nums1, nums2, m, n = nums2, nums1, n, m
	}

	// 初始位置、较短数组长度、中位数所在位置
	imin, imax, halfLen := 0, m, (m+n+1)/2
	for imin <= imax {
		// 第一个数组中位数
		i := (imin + imax) / 2
		j := halfLen - 1
		if i < m && nums2[j-1] > nums1[i] {
			imin = i + 1
		}
	}
}
