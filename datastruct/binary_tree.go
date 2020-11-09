package datastruct

import "container/list"

// 深度优先遍历
// 前序遍历 中左右
// 中序遍历 左中右
// 后续遍历 左右中

// 广度优先遍历(Breadth First Search)，又叫宽度优先搜索或横向优先搜素，是从根基点开始沿着树的高度搜索遍历；
// 广度优先遍历与深度优先遍历的区别
// 广度优先遍历与深度优先遍历的区别在于：广度优先遍历是以层为顺序，将某一层上的所有节点都搜索到了之后才向下一层搜索；
// 而深度优先遍历是将某一条枝上的所有节点都搜索到了之后，才转向搜索另一条枝桠上的所有节点。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 遍历方式 一、迭代  二、递归

// preorderTraversal
// 递归
// 二叉树前序遍历 中左右
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	rs := []int{}
	rs = append(rs, root.Val)
	if root.Left != nil {
		rsLeft := preorderTraversal(root.Left)
		rs = append(rs, rsLeft...)
	}

	if root.Right != nil {
		rsRight := preorderTraversal(root.Right)
		rs = append(rs, rsRight...)
	}

	return rs
}

// preorderTraversalV2
// 迭代方式 前序遍历 二叉树 借助链表
func preorderTraversalV2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	pre := root
	stack := list.New()
	rs := []int{}
	for pre != nil || stack.Len() != 0 {
		// 循环
		for pre != nil {
			rs = append(rs, pre.Val)
			stack.PushBack(pre)
			pre = pre.Left
		}

		if stack.Len() != 0 {
			item := stack.Back()
			pre = item.Value.(*TreeNode)
			pre = pre.Right
			stack.Remove(item)
		}
	}
	return rs
}

// inorderTraversal
// 递归 中序遍历二插树 左中右
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	pre := root
	rs := []int{}
	if pre.Left != nil {
		rs = append(rs, inorderTraversal(pre.Left)...)
	}
	rs = append(rs, pre.Val)
	if pre.Right != nil {
		rs = append(rs, inorderTraversal(pre.Right)...)
	}
	return rs
}

// inorderTraversalV2
// 迭代 中序遍历二叉树 左中右
func inorderTraversalV2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	pre := root
	rs := []int{}

	stack := list.New()

	for pre != nil || stack.Len() != 0 {
		// 现将所有左子树都放到列表中
		for pre != nil {
			stack.PushBack(pre)
			pre = pre.Left
		}

		if stack.Len() != 0 {
			item := stack.Back()
			pre = item.Value.(*TreeNode)
			rs = append(rs, pre.Val)
			pre = pre.Right
			stack.Remove(item)
		}
	}
	return rs
}

// postorderTraversal
// 递归 后续遍历 左右中
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	pre := root
	rs := []int{}
	if pre.Left != nil {
		rs = append(rs, postorderTraversal(pre.Left)...)
	}
	if pre.Right != nil {
		rs = append(rs, postorderTraversal(pre.Right)...)
	}
	rs = append(rs, pre.Val)
	return rs
}

// // 后序遍历-非递归
// func (bt *BinaryTree) PostOrderNoRecursion() []interface{} {
// 	t := bt
// 	stack := list.New()
// 	res := make([]interface{}, 0)
// 	var preVisited *BinaryTree

// 	for t != nil || stack.Len() != 0 {
// 		for t != nil {
// 			stack.PushBack(t)
// 			t = t.Left
// 		}

// 		v := stack.Back()
// 		top := v.Value.(*BinaryTree)

// 		if (top.Left == nil && top.Right == nil) || (top.Right == nil && preVisited == top.Left) || preVisited == top.Right {
// 			res = append(res, top.Data) //visit
// 			preVisited = top
// 			stack.Remove(v)
// 		} else {
// 			t = top.Right
// 		}
// 	}
// 	return res
// }
