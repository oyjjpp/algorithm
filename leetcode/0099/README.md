### 题目
二叉搜索树中的两个节点被错误地交换；  
请在不改变其结构的情况下，恢复这棵树。

#### 案例1
```
输入:[1,3,null,null,2]
{"Val":1,"Left":{"Val":3,"Right":{"Val":2}},}

输出:[3,1,null,null,2] 
{{"Val":3,"Left":{"Val":1,"Right":{"Val":2}},}}
```


#### 案例2
```
输入:[3,1,4,null,null,2]
{"Val":3,"Left":{"Val":1},"Right":{"Val":4,Left:{Val:2}}}

输出:[2,1,4,null,null,3] 
{"Val":2,"Left":{"Val":1},"Right":{"Val":4,Left:{Val:3}}}
```

### 思路


### 参考
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/recover-binary-search-tree/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
