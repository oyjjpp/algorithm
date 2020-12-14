package leetcode

import (
	"testing"
)

func TestMinDistance(t *testing.T) {
	word1, word2 := "horse", "ros"
	rs := minDistance(word1, word2)
	t.Log(rs)

	word1, word2 = "intention", "execution"
	rs = minDistance(word1, word2)
	t.Log(rs)
}
