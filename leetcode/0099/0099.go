package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
// recoverTree
// 恢复一个二叉搜索树
func recoverTree(root *TreeNode)  {
    nums := []int{}
    
    // 定义一个中序遍历函数，将结果保存到数组中
    var inorder func(node *TreeNode)
    inorder = func(node *TreeNode) {
        if node == nil {
            return
        }
        inorder(node.Left)
        nums = append(nums, node.Val)
        inorder(node.Right)
    }
    // 中序遍历二叉树
    inorder(root)
    // 查找交换的节点
    x, y := findTwoSwapped(nums)
    fmt.Println(x, y)
    recover(root, 2, x, y)
}

// findTwoSwapped
// 找到两个交换的索引位置
func findTwoSwapped(nums []int) (int, int) {
    x, y := -1, -1
    for i := 0; i < len(nums) - 1; i++ {
        if nums[i + 1] < nums[i] {
            y = nums[i+1]
            if x == -1 {
                x = nums[i]
            } else {
                break
            }
        }
    }
    return x, y
}

func recover(root *TreeNode, count, x, y int) {
    if root == nil {
        return
    }
    if root.Val == x || root.Val == y {
        if root.Val == x {
            root.Val = y
        } else {
            root.Val = x
        }
        count--
        if count == 0 {
            return
        }
    }
    recover(root.Right, count, x, y)
    recover(root.Left, count, x, y)
}
