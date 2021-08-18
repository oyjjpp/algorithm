### 题目

101 对称二叉树

#### 题目描述

给定一个二叉树，检查它是否是镜像对称的

### 思路

1.根节点为nil则是对称二叉树
2.否则校验左子数和右子数是否相等
3.如果左子数和右子数相等，则校验左子数的左子数和右子数的右子数是否相等 && 左子数的右子数与右子数的左子数是否相等
4.对2.3两步骤进行递归校验

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
链接：<https://leetcode-cn.com/problems/symmetric-tree/>  
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
