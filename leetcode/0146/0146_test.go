package leetcode

import "testing"

func TestLRUCache(t *testing.T) {

	// ["LRUCache","put","put","get","put","get","put","get","get","get"]
	// [[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]
	// [null,null,null,1,null,-1,null,-1,3,4]
	// lruCache := Constructor(2)
	// lruCache.Put(1, 1)
	// lruCache.Put(2, 2)
	// rs := lruCache.Get(1)
	// t.Log(rs)
	// lruCache.Put(3, 3)
	// rs = lruCache.Get(2)
	// t.Log(rs)
	// lruCache.Put(4, 4)
	// rs = lruCache.Get(1)
	// t.Log(rs)
	// rs = lruCache.Get(3)
	// t.Log(rs)
	// rs = lruCache.Get(4)
	// t.Log(rs)

	// ["LRUCache","put","put","put","put","get","get"]
	// [[2],[2,1],[1,1],[2,3],[4,1],[1],[2]]
	// [null,null,null,null,null,-1,3]

	lruCache := Constructor(2)
	lruCache.Put(2, 1)
	lruCache.Put(1, 1)
	lruCache.Println()
	lruCache.Put(2, 3)
	lruCache.Println()
	lruCache.Put(4, 1)
	rs := lruCache.Get(1)
	t.Log(rs)
	rs = lruCache.Get(2)
	t.Log(rs)
	lruCache.Println()
}
