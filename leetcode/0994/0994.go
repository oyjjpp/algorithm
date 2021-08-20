package leetcode

import "container/list"

type Area [2]int

// 地图分析
func orangesRotting(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

	var queue list.List

	// 将所有陆地各加入队列
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				item := Area{i, j}
				queue.PushBack(item)
			}
		}
	}

	// 判断元素的上下左右坐标
	moves := []Area{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}

	// 记录当前遍历的层次
	distince := -1

	for queue.Len() > 0 {
		distince++

		number := queue.Len()

		for i := 0; i < number; i++ {
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
				if inArea(grid, x, y) && grid[x][y] == 1 {
					// 保证不会重复加入
					grid[x][y] = 2
					queue.PushBack(Area{x, y})
				}
			}

		}
	}
	return distince
}

// 判断坐标是否在指定范围内
func inArea(grid [][]int, x, y int) bool {
	return x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0])
}
