package leetcode

// twoSum
// 两数之和
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

// 暴力破解
// 时间复杂多O(n^2)
// 空间复杂度O(1)
func twoSumV2(nums []int, target int) []int {
	rs := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			temp := nums[i] + nums[j]
			if temp == target {
				rs[0] = i
				rs[1] = j
				break
			}
		}
	}
	return rs
}
