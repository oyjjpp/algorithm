package leetcode

import "testing"

func TestLRUCache(t *testing.T) {
	lruCache := Constructor(2)
	lruCache.Put(1, 1)
	lruCache.Put(2, 2)

	// rs := lruCache.Get(1)
	// t.Log(rs)
	// lruCache.Put(3, 3)
	// rs = lruCache.Get(2)
	// t.Log(rs)
}
