package leetcode

import (
	"log"
)

// permute
// 全排列
// @param nums 数字序列
func permute(nums []int) [][]int {
	// 记录结果
	res := [][]int{}
	// 临时使用记录是否已经在路径中
	visited := map[int]bool{}

	// TODO 分析程序
	// index := 0
	// 遍历路径[1,2,3]
	var dfs func(path []int)
	dfs = func(path []int) {
		log.Println(path, visited)
		// TODO 验证程序的逻辑
		// index++
		// log.Println(index, path)
		// 遍历到底层，路径长路与选择列表长度相同
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for _, n := range nums {
			// 校验路径中是否存在
			if is, ok := visited[n]; ok && is {
				continue
			}

			// 选择列表中选取一个元素放入路径中
			path = append(path, n)
			visited[n] = true
			dfs(path)
			log.Println(n)

			// 将该选择从路径中移除
			path = path[:len(path)-1]
			visited[n] = false
		}
	}
	dfs([]int{})
	return res
}
