package leetcode

import (
	"log"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var pre *TreeNode
	isValid := true
	
	var inorder func (root *TreeNode)
	inorder = func(root *TreeNode){
		if root == nil{
			return 
		}
		inorder(root.Left)
		if pre == nil {
			pre = root
		}else{
			if pre.Val>root.Val{
				isValid = false
				return
			}
			pre = root
		}
		log.Println(pre.Val)
		
		inorder(root.Left)
	}
	inorder(root)
	return isValid
}
