package leetcode

import (
	"fmt"
	"math"
)

// K 代表鸡蛋个数
// N 代表楼的高度
// F 临界值
// ? 求F的的最优解
// getMax(helper(K,N-i,mem),helper(K,i-1,mem))
// getMin(res,getMax(helper(K,N-i,mem),helper(K,i-1,mem)))
func superEggDrop(K int, N int) int {
	if K == 1 {
		return N
	}
	if K == 0 {
		return 0
	}
	mem := make(map[string]int)
	res := helper(K, N, &mem)
	return res
}

func helper(K int, N int, mem *map[string]int) int {
	if K == 1 {
		return N
	}
	if N == 0 {
		return 0
	}
	var (
		res  = math.MaxInt32
		temp = fmt.Sprintf("%d,%d", K, N)
		l    = 1
		r    = N
	)
	if (*mem)[temp] != 0 {
		return (*mem)[temp]
	}
	for l <= r {
		i := l + (r-l)>>1
		notbroken := helper(K, N-i, mem)
		broken := helper(K-1, i-1, mem)
		res = getMin(res, getMax(notbroken, broken)+1)
		if notbroken > broken {
			l = i + 1
		} else {
			r = i - 1
		}
	}
	(*mem)[temp] = res
	return res
}

func getMin(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func getMax(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func superEggDropV2(K, N int) int {
	if K < 1 || N < 1 {
		return 0
	}
	//备忘录，存储K个鸡蛋，N层楼条件下的最优化尝试次数
	//cache := [K + 1][N + 1]int{}
	cache := make([][]int, K+1)
	//把备忘录每个元素初始化成最大的尝试次数
	for i := 0; i <= K; i++ {
		cache[i] = make([]int, N+1)
		for j := 1; j <= N; j++ {
			cache[i][j] = j
		}
	}
	for n := 2; n <= K; n++ {
		for m := 1; m <= N; m++ {
			//假设楼层数可以是1---N,
			min := cache[n][m]
			for k := 1; k < m; k++ {
				//M层,N鸡蛋,F（N，K）= Min（Max（ F（N-X，K）+ 1， F（X-1，K-1） + 1）），1<=X<=N
				//(动态规划)
				//鸡蛋碎了
				max := cache[n-1][k-1] + 1
				if cache[n][m-k]+1 > max {
					max = cache[n][m-k] + 1 //鸡蛋没碎
				}
				if max < min {
					min = max
				}
			}
			cache[n][m] = min
		}
	}
	return cache[K][N]
}
