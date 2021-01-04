### 题目

请你为最不经常使用（LFU）缓存算法设计并实现数据结构。

### 题目描述

实现 LFUCache 类：

>LFUCache(int capacity) - 用数据结构的容量 capacity 初始化对象  
>int get(int key) - 如果键存在于缓存中，则获取键的值，否则返回 -1。  
>void put(int key, int value) - 如果键已存在，则变更其值；如果键不存在，请插入键值对。  
>当缓存达到其容量时，则应该在插入新项之前，使最不经常使用的项无效。在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，应该去除最久未使用的键。

### 参考

来源：力扣（LeetCode）
链接：<https://leetcode-cn.com/problems/lfu-cache>
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
