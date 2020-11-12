### 题目
110.平衡二叉树

#### 题目描述
给定一个二叉树，判断它是否是高度平衡的二叉树。

#### 平衡二叉树
一个二叉树每个节点的左右两个子树的高度差的绝对值不超过1。     

### 思路
递归的顺序可以是自顶向下或者自底向上  

#### 自顶向下递归 
前序遍历方式   
1.判断节点是否为nil，如果是nil则是平衡二叉树  
2.判断左右子数的高度差是否大于一，如果大于一，则不是平衡二叉树，否则继续校验子树  
3.判断左子数是否为平衡二叉树  
4.判断右子数是否平衡二叉树  

#### 自底向上
后续遍历  
1.求左，右子树高度
2.左子数，右子数高度是否为-1, 左右高度差大于1，则返回-1
3.当前树的高度是否大于等于0

### 示例
```golang
data := &TreeNode{
	Val:3,
	Left:&TreeNode{
		Val:9,
	},
	Right:&TreeNode{
		Val:20,
		Left:&TreeNode{Val:15},
		Right:&TreeNode{Val:7},
	},
}

data := &TreeNode{
	Val:1,
	Left:&TreeNode{
		Val:2,
		Left:&TreeNode{
			Val:3,
			Left:&TreeNode{Val:4},
			Right:&TreeNode{Val:4},
		},
		Right:&TreeNode{Val:3},
	},
	Right:&TreeNode{
		Val:2,
	},
}

```

### 代码
```golang

```

### 参考
来源：力扣（LeetCode）  
链接：https://leetcode-cn.com/problems/balanced-binary-tree/  
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
