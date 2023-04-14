package hot100

import "log"

/*
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码还未经过力扣测试，仅供参考，如有疑惑，可以参照我写的 java 代码对比查看。

// 计算从起点 start 到终点 target 的最近距离
func BFS(start *Node, target *Node) int {
    q := make([]*Node, 0) // 核心数据结构
    visited := make(map[*Node]bool) // 避免走回头路

    q = append(q, start) // 将起点加入队列
    visited[start] = true
    step := 0 // 记录扩散的步数

    for len(q) > 0 {
        sz := len(q)
        /// 将当前队列中的所有节点向四周扩散
        for i := 0; i < sz; i++ {
            cur := q[0]
            q = q[1:]
            // 划重点：这里判断是否到达终点
            if cur == target {
                return step
            }
            // 将 cur 的相邻节点加入队列
            for _, x := range cur.adj() {
                if !visited[x] {
                    q = append(q, x)
                    visited[x] = true
                }
            }
        }
        // 划重点：更新步数在这里
        step++
    }
}
*/

// 111. 二叉树的最小深度
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	minPath := 1

	for len(queue) > 0 {
		sz := len(queue)

		for i := 0; i < sz; i++ {
			node := queue[i]
			log.Println(i, node.Val)
			// 碰到叶子节点
			if node.Left == nil && node.Right == nil {
				return minPath
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[sz:]
		minPath++
	}
	return minPath
}

// 752. 打开转盘锁
func openLock(deadends []string, target string) int {
	var plusOne func(s string, j int) string
	plusOne = func(s string, j int) string {
		ch := []byte(s)
		if ch[j] == '9' {
			ch[j] = '0'
		} else {
			ch[j] += 1
		}
		return string(ch)
	}

	var minusOne func(s string, j int) string
	minusOne = func(s string, j int) string {
		ch := []byte(s)
		if ch[j] == '0' {
			ch[j] = '9'
		} else {
			ch[j] -= 1
		}
		return string(ch)
	}
	// 死亡密码
	deads := make(map[string]bool)
	for _, s := range deadends {
		deads[s] = true
	}

	// 记录已经穷举过的密码，防止走回头路
	visited := make(map[string]bool)

	queue := make([]string, 0)
	queue = append(queue, "0000")
	visited["0000"] = true

	step := 0

	for len(queue) > 0 {
		sz := len(queue)

		for i := 0; i < sz; i++ {
			node := queue[0]
			queue = queue[1:]

			/* 判断是否到达终点 */
			if _, ok := deads[node]; ok {
				continue
			}
			if node == target {
				return step
			}

			for j := 0; j < 4; j++ {
				up := plusOne(node, j)
				if _, ok := visited[up]; !ok {
					queue = append(queue, up)
					visited[up] = true
				}
				down := minusOne(node, j)
				if _, ok := visited[down]; !ok {
					queue = append(queue, down)
					visited[down] = true
				}
			}
		}
		step++
	}
	return -1
}

// 向上拨动
func plusOne(s string, j int) string {
	ch := []byte(s)
	if ch[j] == '9' {
		ch[j] = '0'
	} else {
		ch[j] += 1
	}
	return string(ch)
}

// 向下拨动
// 将 s[i] 向下拨动一次
func minusOne(s string, j int) string {
	ch := []byte(s)
	if ch[j] == '0' {
		ch[j] = '9'
	} else {
		ch[j] -= 1
	}
	return string(ch)
}
