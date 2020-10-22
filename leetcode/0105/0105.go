package leetcode

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// preorder 前序遍历参数  中左右
// inorder 中序遍历参数	左中右
func buildTree(preorder []int, inorder []int) *TreeNode {
    p, d := 0, make(map[int]int, len(inorder))
    for i, j := range inorder {
        d[j] = i
    }
    var f func(int, int) *TreeNode; f = func(i, j int) (t *TreeNode) {
        if i < j {
            t = &TreeNode{Val: preorder[p]}
            p++
            t.Left = f(i, d[t.Val])
            t.Right = f(d[t.Val] + 1, j)
        }
        return 
    }
    return f(0, len(inorder))
}
