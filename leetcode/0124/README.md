### 题目
给定一个非空二叉树，返回其最大路径和。  
本题中，路径被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。

#### 案例1
```
输入: [1,2,3]
参数结构：
data := &TreeNode{
	Val:   1,
	Left:  &TreeNode{Val: 2},
	Right: &TreeNode{Val: 3},
}
输出: 6 
```

#### 案例2
```
输入: [-10,9,20,null,null,15,7]
参数结构：
data1 := &TreeNode{
	Val:  -10,
	Left: &TreeNode{Val: 9},
	Right: &TreeNode{
		Val:   20,
		Left:  &TreeNode{Val: 15},
		Right: &TreeNode{Val: 7},
	},
}
输出: 42
```

### 思路
1. 使用一个局部变量记录当前子树路径之和最大值
2. 后续遍历二叉树每个子树的和
3. 当前值与最大值做比较

### 参考
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-maximum-path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
