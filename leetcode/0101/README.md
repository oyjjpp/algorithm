### 题目
 对称二叉树

#### 题目描述
给定一个二叉树，检查它是否是镜像对称的      

### 思路
1.中序遍历二叉树  
2.结果保存到一个数组中  
3.比较数组前后索引是否一致  

### 示例
```golang
data := &TreeNode{
	Val:1,
	Left:&TreeNode{
		Val:2,
		Left:&TreeNode{Val:3},
		Right:&TreeNode{Val:4},
	},
	Right:&TreeNode{
		Val:2,
		Left:&TreeNode{Val:4},
		Right:&TreeNode{Val:3},
	},
}

3241423
返回
```

### 代码
```golang

```

### 参考
来源：力扣（LeetCode）  
链接：https://leetcode-cn.com/problems/symmetric-tree/  
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
