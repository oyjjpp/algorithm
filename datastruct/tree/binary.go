package tree

type Node struct {
	val   int
	left  *Node
	right *Node
}

// pushOne
// 二叉树所有节点加一
func pushOne(root *Node) {
	if root == nil {
		return
	}
	root.val += 1
	pushOne(root.left)
	pushOne(root.right)
}

// isSameThee
// 判断两个二叉树是否完全相同
func isSameThee(root1, root2 *Node) bool {
	// 都为空的，则相同
	if root1 == nil && root2 == nil {
		return true
	}

	// 一个为空，一个非空
	if root1 == nil && root2 != nil {
		return false
	}
	if root2 == nil && root1 != nil {
		return false
	}

	// 两个都非空，但val不一样
	if root1.val != root2.val {
		return false
	}
	return isSameThee(root1.left, root2.left) && isSameThee(root1.right, root2.right)
}

// invertBT
// 通过递归方式翻转一颗二叉树
func invertBT(root *Node) *Node {
	if root == nil {
		return root
	}

	// 交换左右孩子节点
	root.left, root.right = root.right, root.left

	// 递归处理左右子树
	invertBT(root.left)
	invertBT(root.right)
	return root
}
