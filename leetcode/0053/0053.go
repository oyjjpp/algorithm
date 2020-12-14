package leetcode

// maxSubArray
// 连续子数组的最大和
func maxSubArray(nums []int) int {
	if len(nums) == 0 || nums == nil {
		return 0
	}

	dp := make([]int, len(nums))
	// base case
	dp[0] = nums[0]

	maxNums := dp[0]
	// 通过DP table求每个位置连续子数组的和
	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
		maxNums = max(maxNums, dp[i])
	}

	/*
	   maxNums := -1 << 31
	   for i:=0; i < len(dp);i++{
	       maxNums = max(maxNums, dp[i])
	   }
	*/
	return maxNums
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
