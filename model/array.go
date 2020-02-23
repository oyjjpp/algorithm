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
