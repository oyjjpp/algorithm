### 题目
965. 单值二叉树

#### 题目描述
如果二叉树每个节点都具有相同的值，那么该二叉树就是单值二叉树。  
只有给定的树是单值二叉树时，才返回 true；否则返回 false。    

### 思路
1.深度遍历二叉树所有节点  
2.节点存储到map中  
3.如果map长度为1则位单值二叉树，否则不是单值二叉树  

### 示例
```golang
二叉树
data := &TreeNode{
	Val:1,
	Left:&TreeNode{
		Val:1,
	},
	Right:&TreeNode{
		Val:2,
	},
}

返回false
```

### 代码
```golang

```

### 参考
来源：力扣（LeetCode）  
链接：https://leetcode-cn.com/problems/univalued-binary-tree/  
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
