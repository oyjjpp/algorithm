package leetcode

type TreeNode struct{
    Val int
    Left *TreeNode
    Right *TreeNode
}

// minDepth
// 二叉树最小深度
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
}

// minDepth
// 二叉树最小深度
// DFS（深度优先搜索）
func minDepthDFS(root *TreeNode) int {
    if root == nil {
        return 0
    }
    if root.Left == nil && root.Right == nil {
        return 1
    }
    minD := math.MaxInt32
    if root.Left != nil {
        minD = min(minDepth(root.Left), minD)
    }
    if root.Right != nil {
        minD = min(minDepth(root.Right), minD)
    }
    return minD + 1
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

// minDepth
// 二叉树最小深度
// BFS（广度优先搜索） 
func minDepthBFS(root *TreeNode) int {
    if root == nil {
        return 0
    }
    // 使用切片定义一个存储树结构的队列
    queue := []*TreeNode{}
    // 记录当前状态
    count := []int{}
    
    queue = append(queue, root)
    // root 本身就是一层，depth 初始化为 1
    count = append(count, 1)
    
    // 将当前队列中的所有节点向四周扩散
    for i := 0; i < len(queue); i++ {
        // 当前节点
        node := queue[i]
        // 深度
        depth := count[i]
        // 判断是否达到终点
        // 如果是叶子节点，则直接返回深度
        if node.Left == nil && node.Right == nil {
            return depth
        }
        
        // 将相邻的节点加入到队列中
        // 校验左子数是否为空
        if node.Left != nil {
            queue = append(queue, node.Left)
            count = append(count, depth + 1)
        }
        // 校验右子数是否为空
        if node.Right != nil {
            queue = append(queue, node.Right)
            count = append(count, depth + 1)
        }
    }
    return 0
}
