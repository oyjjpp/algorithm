#

## 二叉树

### 题目

（108. ）将有序数组转换为二叉搜索树

#### 题目描述

将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树；  
本题中，一个高度平衡二叉树是指一个二叉树每个节点的左右两个子树的高度差的绝对值不超过1。  

#### 平衡二叉树

一个二叉树每个节点的左右两个子树的高度差的绝对值不超过1。

#### 二叉搜索树

二叉搜素树中序遍历结果是递增顺序

### 思路

本题目中提供的几条重要线索  
1、参数为有序的数组  
2、转换结果为二叉搜索树  
3、高度平衡的二叉搜索树  

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
链接：<https://leetcode-cn.com/problems/balanced-binary-tree/>  
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
