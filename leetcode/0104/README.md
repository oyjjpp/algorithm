### 题目
104. 二叉树的最大深度

#### 题目描述
给定一个二叉树，找出其最大深度。  

#### 二叉树的深度
二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。  

#### 说明
叶子节点是指没有子节点的节点。     

### 思路


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
