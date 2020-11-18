package leetcode

import(
    "log"
)

// permute
// 全排列
// @param nums 数字序列
func permute(nums []int) [][]int {
	res := [][]int{}
	visited := map[int]bool{}

    // 遍历路径[1,2,3]
	var dfs func(path []int)
	dfs = func(path []int) {
        // 遍历到底层，路径长路与选择列表长度相同
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
        log.Println(path, visited)
		for _, n := range nums {
            // 校验路径中是否存在
			if is ,ok := visited[n];ok && is {
				continue
			}
            // 选择列表中选取一个元素放入路径中
			path = append(path, n)
			visited[n] = true
			dfs(path)
            // 将该选择从路径中移除
			path = path[:len(path)-1]
			visited[n] = false
		}
	}

	dfs([]int{})
	return res
}
