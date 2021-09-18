package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// buildTree
// 构建二叉树
// @param preorder 前序遍历参数  中左右
// @param inorder 中序遍历参数	左中右
func buildTree(preorder []int, inorder []int) *TreeNode {
	// 节点为0
	if len(preorder) == 0 {
		return nil
	}
	// 前序遍历第一个元素为根节点
	root := &TreeNode{preorder[0], nil, nil}

	// 在中序遍历中查找根节点位置
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}

	// 递归构建左右子树
	// 参数为左子数的前序遍历，中序遍历
	// 前序遍历的左子树需要通过中序遍历计算出的长度确定
	index := len(inorder[:i]) + 1
	root.Left = buildTree(preorder[1:index], inorder[:i])

	// 参数为右子数的前序遍历，中序遍历
	root.Right = buildTree(preorder[index:], inorder[i+1:])
	return root
}
