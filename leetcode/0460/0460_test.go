package leetcode

import "testing"

func TestLFUCache(t *testing.T) {

	// 	输入：
	// ["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get", "get"]
	// [[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]
	// 输出：
	// [null, null, null, 1, null, -1, 3, null, -1, 3, 4]

	// 解释：
	// LFUCache lFUCache = new LFUCache(2);
	// lFUCache.put(1, 1);
	// lFUCache.put(2, 2);
	// lFUCache.get(1);      // 返回 1
	// lFUCache.put(3, 3);   // 去除键 2
	// lFUCache.get(2);      // 返回 -1（未找到）
	// lFUCache.get(3);      // 返回 3
	// lFUCache.put(4, 4);   // 去除键 1
	// lFUCache.get(1);      // 返回 -1（未找到）
	// lFUCache.get(3);      // 返回 3
	// lFUCache.get(4);      // 返回 4

	lfuCache := Constructor(2)
	lfuCache.Put(1, 1)
	lfuCache.Put(2, 2)
	rs := lfuCache.Get(1)
	t.Log(rs)
	lfuCache.Put(3, 3)
	rs = lfuCache.Get(2)
	t.Log(rs)
	rs = lfuCache.Get(3)
	t.Log(rs)
	lfuCache.Put(4, 4)
	rs = lfuCache.Get(1)
	t.Log(rs)
	rs = lfuCache.Get(3)
	t.Log(rs)
	rs = lfuCache.Get(4)
	t.Log(rs)
}
