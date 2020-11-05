package leetcode

import(
	"log"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// recoverTree
// 恢复一个二叉搜索树
func recoverTree(root *TreeNode)  {
	// 定义一个数组保存中序遍历结果
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
    log.Println(nums)
    // 查找交换的节点
    x, y := findTwoSwapped(nums)
    // TODO 如果都是-1 则代表已经是二叉搜索树 1237564
    log.Println(x, y)
    recover(root, 2, x, y)
}

// 1237564
// findTwoSwapped
// 找到两个交换的索引位置值
func findTwoSwapped(nums []int) (int, int) {
	// 用于存储错误交换的位置
    x, y := -1, -1
    for i := 0; i < len(nums) - 1; i++ {
		// TODO 两个错误的位置一定是相连的吗？
		// 当后一个小于前一个的时候进行存储
        if nums[i + 1] < nums[i] {
			// 为什么y可以进行替换？找到最终的错的位置
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

// 恢复二叉搜索树
// recover
// @param root 二叉树
func recover(root *TreeNode, count, x, y int) {
    if root == nil {
        return
    }
    // 二叉树中查找对应的值然后进行替换
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
    recover(root.Left, count, x, y)
    recover(root.Right, count, x, y)
}

// recoverTreeV2
// 恢复一个二叉搜索树 通过节点指针保存
func recoverTreeV2(root *TreeNode)  {
	var x, y, pre  *TreeNode
    
    // 定义一个中序遍历函数，将结果保存到数组中
    var inorder func(node *TreeNode)
    inorder = func(node *TreeNode) {
        if node == nil {
            return
        }
        log.Println("node", node.Val)
        inorder(node.Left)
        if pre==nil {
			pre = node
			log.Println(pre.Val)
		} else {
			if pre.Val > node.Val {
				y = node
				if x==nil {
					x = pre
				}
			}
			// 保存当前节点作为上一个节点
			pre = node
		}
        inorder(node.Right)
    }
    // 中序遍历二叉树
    inorder(root)
    if x!= nil && y != nil{
		recover(root, 2, x.Val, y.Val)
	}
}
