package leetcode

// maxCoins
// 312. 戳气球
// 使用回溯算法——暴力穷举
func maxCoins(nums []int) int {
	return 0
	/*
		res := -1 << 31

		var backtrack func(nums []int, socre int)

		// 临时使用记录是否已经在路径中
		// visited := map[int]bool{}

		// backtrack
		// 回溯算法
		// @param nums 一组气球
		// @param score 当前计算得分数
		backtrack = func(nums []int, score int) {
			// 所有气球都被戳破
			if nums == nil || len(nums) == 0 {
				res = max(res, score)
				return
			}

			for i := 0; i < len(nums); i++ {
				// 当前分数
				// 越界检查
				left := 1
				if i > 0 {
					left = nums[i-1]
				}
				right := 1
				if i < len(nums)-1 {
					right = nums[i+1]
				}
				// 硬币个数
				point := left * nums[i] * right

				temp := nums[i]

				// 做选择
				// 在nums中删除元素nums[i]
				// 递归回溯
				backtrack(nums, score+point)

				// 撤销选择
				// 将temp还原nums[i]
			}
		}
		backtrack(nums, 0)
		return res
	*/
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// maxCoinsV2
// 312. 戳气球
// 使用动态规划
func maxCoinsV2(nums []int) int {
	n := len(nums)

	// 添加两侧得虚拟气球
	// points[0] = points[n+1] = 1
	points := make([]int, n+2)
	points[0], points[n+1] = 1, 1

	for i := 1; i <= n; i++ {
		points[i] = nums[i-1]
	}

	// base case
	dp := makeSlice(n+2, n+2)

	// 状态转移
	// i应该从下到上
	for i := n; i >= 0; i-- {
		// j应该从左到右
		for j := i + 1; j < n+2; j++ {
			// 最后戳破得气球是哪个？
			for k := i + 1; k < j; k++ {
				// 择优选择
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j]+points[i]*points[j]*points[k])
			}
		}
	}
	return dp[0][n+1]
}

func makeSlice(m, n int) [][]int {
	data := make([][]int, m)
	for key := range data {
		data[key] = make([]int, n)
	}
	return data
}
