### 题目

根据一棵树的前序遍历与中序遍历构造二叉树。

### 注意

你可以假设树中没有重复的元素

#### 案例1

```
输入:
前序遍历：[3, 9, 20, 15, 7]
中序遍历：[9, 3, 15, 20, 7]
输出:
{"Val":3,"Left":{"Val":9,"Left":null,"Right":null},"Right":{"Val":20,"Left":{"Val":15,"Left":null,"Right":null},"Right":{"Val":7,"Left":null,"Right":null}}}
```

### 思路

1. 先序遍历的顺序是根节点，左子树，右子树；中序遍历的顺序是左子树，根节点，右子树；  
2. 需要根据先序遍历得到根节点，然后在中序遍历中找到“根节点”的位置；  
3. “根节点”的左边就是左子树的节点，右边就是右子树的节点；  
4. 生成左子树和右子树就可以递归的进行了。  

### 参考

来源：力扣（LeetCode）
链接：<https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/>
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
