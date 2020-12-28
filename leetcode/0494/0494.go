package leetcode

import "fmt"

// findTargetSumWays
// 494. 目标和
// 回溯算法——暴力求解
func findTargetSumWays(nums []int, S int) int {
	if len(nums) == 0 {
		return 0
	}
	res := 0
	var backtrack func(nums []int, i, rest int)
	// backtrack
	// 回溯算法
	// @param nums 所提供得数组
	// @param i 当前索引
	// @param rest taget 剩余的值
	backtrack = func(nums []int, i, rest int) {
		// base case
		if i == len(nums) {
			if rest == 0 {
				res++
			}
			return
		}
		// 做出选择
		backtrack(nums, i+1, rest-nums[i])
		backtrack(nums, i+1, rest+nums[i])
	}
	backtrack(nums, 0, S)
	return res
}

// findTargetSumWays
// 494. 目标和
// 动态规划——解决重叠子问题提升效率
func findTargetSumWaysV2(nums []int, S int) int {
	if len(nums) == 0 {
		return 0
	}

	// 备忘录，解决重叠子问题
	meno := map[string]int{}

	var dp func(nums []int, i, rest int) int
	// backtrack
	// 动态规划
	// @param nums 所提供得数组
	// @param i 当前索引
	// @param rest taget 剩余的值
	dp = func(nums []int, i, rest int) int {
		// base case
		if i == len(nums) {
			if rest == 0 {
				return 1
			}
			return 0
		}
		key := fmt.Sprintf("%d,%d", i, rest)
		if rs, ok := meno[key]; ok {
			return rs
		}

		res := dp(nums, i+1, rest-nums[i]) + dp(nums, i+1, rest+nums[i])
		meno[key] = res
		return res
	}
	return dp(nums, 0, S)
}
