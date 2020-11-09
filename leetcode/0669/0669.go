package leetcode

import (
	"encoding/json"
	"log"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// trimBST
// 修剪二叉搜索树
// @param root 二叉树
// @param low 最小边界
// @param high 最大边界
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return root
	}
	/*
		data := &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 1},
			Right: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 6},
			},
		}
	*/
	// [3, 5]
	// 当前节点大于最大边界，则当前节点及右节点已经越界，所以只保留左子数
	if root.Val > high {
		root = trimBST(root.Left, low, high)
		return root
	}
	// 当前节点小于最小边界，则当前节点及左节点已经越界，所以值保留右子数
	if root.Val < low {
		root = trimBST(root.Right, low, high)
		rs, _ := json.Marshal(root)
		log.Println(string(rs))
		return root
	}
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)

	return root
}
