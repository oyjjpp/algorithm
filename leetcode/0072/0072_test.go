package leetcode

import (
	"testing"
)

func TestMinDistance(t *testing.T) {
	word1, word2 := "horse", "ros"
	rs := minDistance(word1, word2)
	t.Log(rs)

	// word1, word2 = "intention", "execution"
	// rs = minDistance(word1, word2)
	// t.Log(rs)
    
    // word1, word2 = "algorithm", "altruistic"
    // rs = minDistance(word1, word2)
	// t.Log(rs)
    
    word1, word2 = "dinitrophenylhydrazine", "benzalphenylhydrazone"
    rs = minDistance(word1, word2)
	t.Log(rs)
}

func TestMinDistanceV(t *testing.T) {
	word1, word2 := "horse", "ros"
	rs := minDistanceV(word1, word2)
	t.Log(rs)

	word1, word2 = "intention", "execution"
	rs = minDistance(word1, word2)
	t.Log(rs)
    
    word1, word2 = "algorithm", "altruistic"
    rs = minDistance(word1, word2)
	t.Log(rs)
    
    word1, word2 = "dinitrophenylhydrazine", "benzalphenylhydrazone"
    rs = minDistanceV(word1, word2)
	t.Log(rs)
}

func TestCreateSlice(t *testing.T){
    data := createSlice(3, 2)
    t.Log(data)
}
