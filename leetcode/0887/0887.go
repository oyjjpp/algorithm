package leetcode

import (
	"fmt"
)

// superEggDrop
// 887. 鸡蛋掉落
func superEggDrop(K int, N int) int {
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
			// log.Println(rs)
			return rs
		}

		res := 1<<31 - 1
		for i := 1; i <= N; i++ {
			// 最坏情况
			// 鸡蛋没碎：K->K, 0~N -> i+1 ~ N
			// 鸡蛋碎了：K->K-1, 0~N -> 0 ~i-1
			maxNums := max(dp(K, N-i), dp(K-1, i-1)) + 1
			// 最少扔鸡蛋次数
			res = min(res, maxNums)
		}
		meno[key] = res
		return res
	}
	return dp(K, N)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// superEggDrop
// 887. 鸡蛋掉落
// 使用二分法提升效率
func superEggDropV2(K int, N int) int {
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
			// 碎了
			broken := dp(K-1, mid-1)
			// 没碎
			notBroken := dp(K, N-mid)

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
