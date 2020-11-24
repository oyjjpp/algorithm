### 题目

（111.）二叉树的最小深度

#### 题目描述

给定一个二叉树，找出其最小深度。  
最小深度是从根节点到最近叶子节点的最短路径上的节点数量。  

#### 说明
叶子节点是指没有子节点的节点。

### 示例

#### 示例一
```conf
输入：root = [3,9,20,null,null,15,7]
输出：2
```

#### 示例二
```conf
输入：root = [2,null,3,null,4,null,5,null,6]
输出：5
```

### 思路
明确起点和终点  
起点：二叉树的根节点  
终点：最靠近根节点的叶子页面  

叶子节点就是两个子节点都是nil的节点  
```golang
if node.Left == nil && node.Right==nil {
	// 叶子节点
}
```


### 代码

```golang

```

### 参考

来源：力扣（LeetCode）  
链接：<https://leetcode-cn.com/problems/minimum-depth-of-binary-tree//>  
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。  

[labuladong的算法小抄](https://labuladong.gitbook.io/algo/di-ling-zhang-bi-du-xi-lie/bfs-kuang-jia)
