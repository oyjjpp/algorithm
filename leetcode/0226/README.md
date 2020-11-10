### 题目
226. 翻转二叉树

#### 题目描述
翻转一棵二叉树。      

### 思路
1.前序遍历二叉树  
2.定义一个临时变量用于存储临时二叉树节点  
3.左右子数进行交换  

### 示例
```golang
data := &TreeNode{
	Val:4,
	Left:&TreeNode{
		Val:2,
		Left:&TreeNode{Val:1},
		Right:&TreeNode{Val:3},
	},
	Right:&TreeNode{
		Val:7,
		Left:&TreeNode{Val:7},
		Right:&TreeNode{Val:9},
	},
}

返回
{"Val":4,"Left":{"Val":7,"Left":{"Val":9,"Left":null,"Right":null},"Right":{"Val":7,"Left":null,"Right":null}},"Right":{"Val":2,"Left":{"Val":3,"Left":null,"Right":null},"Right":{"Val":1,"Left":null,"Right":null}}}
```

### 代码
```golang

```

### 参考
来源：力扣（LeetCode）  
链接：https://leetcode-cn.com/problems/invert-binary-tree/  
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
