package tree

import "log"

// isValidBST
// 是否有有效二叉搜索树
func isValidBST(root *Node) bool {
	var pre *Node
	isValid := true

	// 通过中序遍历的有序性，校验是否为有效二叉树
	var inorder func(root *Node)
	inorder = func(root *Node) {
		if root == nil {
			return
		}
		inorder(root.left)
		if pre == nil {
			pre = root
		} else {
			if pre.val >= root.val {
				isValid = false
				return
			}
			pre = root
		}
		log.Println(pre.val)
		inorder(root.right)
	}
	inorder(root)
	return isValid
}

// isInBST
// 校验目标值是否在二叉搜索树中
func isInBST(root *Node, target int) bool {
	if root == nil {
		return false
	}
	if root.val == target {
		return true
	} else if root.val < target {
		return isInBST(root.right, target)
	} else {
		return isInBST(root.left, target)
	}
}

// insertIntoBST
// 在BST中插入一个数
// 一旦涉及“改”，函数就要返回Node类型，并且对递归调用的返回值进行接收
func insertIntoBST(root *Node, val int) *Node {
	if root == nil {
		return &Node{val: val}
	}
	// 如果已存在，则不再重复擦汗如，直接返回
	if root.val == val {
		return root
	}
	// val小， 则应该插入到左子树
	if root.val > val {
		root.left = insertIntoBST(root.left, val)
	}
	// val大，则应该插入到右子树
	if root.val < val {
		root.right = insertIntoBST(root.right, val)
	}
	return root
}
