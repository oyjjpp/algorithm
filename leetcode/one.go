package leetcode

// 暴力破解
// 时间复杂多O(n^2)
// 空间复杂度O(1)
func twoSum(nums []int, target int) []int {
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

// 通过map特性解决
// 时间复杂O(N)
// 空间复杂O(N)
func twoSumV2(nums []int, target int) []int {
	rs := map[int]int{}
	for k, v := range nums {
		if value, ok := rs[target-v]; ok {
			return []int{value, k}
		}
		rs[v] = k
	}
	return nil
}
