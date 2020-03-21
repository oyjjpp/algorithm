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

// 分治法
// findMedianSortedArrays2 给出两个有序数组，假设两个数组的长度和是len，
// 如果len为奇数，那么我们求的就是两个数组合并后的第 (len >> 1) + 1 大的数，
// 如果len为偶数，就是第 (len >> 1) 和 (len >> 1) + 1 两个数的平均数
// 给定两个有序数组，求第k大数,如果我们从A和B中分别取前k/2个元素，其中必然有一部分是是在数组C的前k个数里
// 设mid=k/2，当A[mid-1]<B[mid-1]时，可以断定A的前mid个元素是在C的前k个数里
// 时间复杂度是 O(log(m+n))
func findMedianSortedArraysV2(nums1 []int, nums2 []int) float64 {
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

// max
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// min
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*

a.
我们重新看题目，要找中位数，就是要找第k大的数（k = (L/2 + 1)，其中L是上面提到的合并后新数组的长度，
当L是偶数时，要求第(L/2)和第(L/2+1)的两个数）。当我们舍弃掉一部分，假设舍弃部分的长度为length，
那么接下来就是在剩下的数组里求第(k-length)大的数,逐层缩小范围，直到两数组其中一个走完，或者要求的是第1大的元素，
就可以直接返回结果了。


b. 那如何“选择”要舍弃哪部分呢？既然是要找合并后的数组C的第k大元素，即 C[k-1]，
那如果我们从A和B中分别取前k/2个元素，其中必然有一部分是是在数组 C 的前 k 个数里。
设 mid = k / 2，当 A[mid - 1] < B[mid - 1] 时，可以断定 A 的前 mid 个元素是在 C 的前 k 个数里（此处可用反证法得证），
那么我们则舍弃 A 的前 mid 个元素。反之则舍弃 B 的前 mid 个元素。现在数组 A 或者 B 已经舍弃掉 k/2 个元素，缩小查找范围了，
那接下来可以按照同样的方法继续选择吗？当然！现在剩下总共 (L - mid) 个元素，且 A 和 B 依旧有序，要找的是第 (k - mid) 大的元素，
所以我们可以按照上面的方法继续递归选择下去，直到找到目标元素！


c. 复杂度分析：每次从合并后数组 C 里减少 k/2 个元素，直到找到目标元素。所以时间复杂度是 O(log L) = O(log (m + n)) ！


*/
