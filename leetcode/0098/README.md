### 题目
给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：
>节点的左子树只包含小于当前节点的数。  
>节点的右子树只包含大于当前节点的数。  
>所有左子树和右子树自身必须也是二叉搜索树。  

### 思路
通过中序遍历，然后对比前后节点的值，如果前一个节点值大于当前节点值，则不是有效二叉搜索树

### 代码
```golang
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
			if pre.Val>=root.Val{
				isValid = false
				return
			}
			pre = root
		}
		log.Println(pre.Val)
		inorder(root.Right)
	}
	inorder(root)
	return isValid
}
```

### 参考
来源：力扣（LeetCode）  
链接：https://leetcode-cn.com/problems/validate-binary-search-tree  
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
