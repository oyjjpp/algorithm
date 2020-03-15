package model

// 两数之和
// 给定一个已按照升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。
// 函数应该返回这两个下标值 index1 和 index2，其中 index1 必须小于 index2。
// 解题思路：循环相加 碰到相等的则返回
// 时间复杂度O(n^2)
// 空间复杂度O(1)
// 暴力破解方法
func TwoSum(numbers []int, target int) []int {
	rs := make([]int, 2)
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				return []int{i, j}
			}
		}
	}
	return rs
}

// 哈希方式解决两数之和
func TwoSumHash(numbers []int, target int) []int {
	data := make(map[int]int, len(numbers))

	for k, v := range numbers {
		sub := target - v
		if j, ok := data[sub]; ok {
			return []int{j, k}
		} else {
			data[v] = k
		}
	}
	return nil
}

// majorityElement
// 多数元素
// 给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
func majorityElement(nums []int) int {
	rs := make(map[int]int)
	for _, k := range nums {
		if _, ok := rs[k]; ok {
			rs[k]++
		} else {
			rs[k] = 1
		}
	}

	length := len(nums)
	for k, v := range rs {
		if v > length/2 {
			return k
		}
	}
	return 0
}

// singleNumber
// 暴力破解
// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
func singleNumber(nums []int) int {
	var rs bool
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			if nums[i] == nums[j] {
				rs = true
			}
		}
		if rs == false {
			return nums[i]
		}
		rs = false
	}
	return 0
}

// singleNumberV2
// 异或 方式计算
func singleNumberV2(nums []int) int {
	var rs int
	for _, v := range nums {
		rs = rs ^ v
	}
	return rs
}

// searchMatrix
// 暴力破解
// 编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：
func searchMatrix(matrix [][]int, target int) bool {
	for _, v := range matrix {
		for _, item := range v {
			if item == target {
				return true
			}
			if item > target {
				break
			}
		}
	}
	return false
}

// searchMatrixV2
// 向量
// 编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：
func searchMatrixV2(matrix [][]int, target int) bool {
	// 行
	row := len(matrix) - 1
	// 列
	col := 0

	for row >= 0 && col < len(matrix[0]) {
		if matrix[row][col] > target {
			row--
		} else if matrix[row][col] < target {
			col++
		} else { // found it
			return true
		}
	}

	return false
}

// merge
// 合并两个有序数组
func merge(nums1 []int, m int, nums2 []int, n int) {
	indexm, indexn := m-1, n-1
	for i := m + n - 1; i >= 0; i-- {
		// 无nums1元素 将nums2赋值nums1
		if indexm < 0 {
			nums1[i] = nums2[indexn]
			indexn--
			continue
		}
		// 无nums2元素，将nums1赋值nums1
		if indexn < 0 {
			nums1[i] = nums1[indexm]
			indexm--
			continue
		}

		// 都存在时 哪个大哪个先赋值nums1
		if nums1[indexm] >= nums2[indexn] {
			nums1[i] = nums1[indexm]
			indexm--
			continue
		}
		if nums1[indexm] < nums2[indexn] {
			nums1[i] = nums2[indexn]
			indexn--
			continue
		}
	}
}
