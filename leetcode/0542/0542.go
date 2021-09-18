package leetcode

import "container/list"

// 每个格子是mat中对应位置元素到最近0的距离
// 连个相邻元素间的距离为1
func updateMatrix(mat [][]int) [][]int {
	m := len(mat)
	if m == 0 {
		return nil
	}
	n := len(mat[0])
	var queue list.List

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				item := Area{i, j}
				queue.PushBack(item)
			} else {
				mat[i][j] = -1
			}
		}
	}

	// 判断元素的上下左右坐标
	moves := []Area{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}

	for queue.Len() > 0 {
		// 队列中获取一个单元格
		element := queue.Front()
		queue.Remove(element)
		if element == nil {
			continue
		}

		node, ok := element.Value.(Area)
		if !ok {
			continue
		}

		r := node[0]
		c := node[1]
		// 判断当前位置的周围（上下左右）
		for _, value := range moves {
			x := r + value[0]
			y := c + value[1]
			// 判断是否越界
			// 判断是否为海洋
			if inArea(mat, x, y) && mat[x][y] == -1 {
				// 保证不会重复加入
				mat[x][y] = mat[r][c] + 1
				queue.PushBack(Area{x, y})
			}
		}
	}

	return mat
}

type Area [2]int

// 判断坐标是否在指定范围内
func inArea(grid [][]int, x, y int) bool {
	return x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0])
}
