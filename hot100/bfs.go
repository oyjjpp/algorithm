package hot100

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
