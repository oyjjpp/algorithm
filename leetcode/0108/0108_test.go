package leetcode

import (
	"encoding/json"
	"log"
	"testing"
)

func TestSortedArrayToBST(t *testing.T) {
	data := []int{-10, -3, 0, 5, 9}
	root := sortedArrayToBST(data)
	if rs, err := json.Marshal(root); err == nil {
		log.Println(string(rs))
	}
}
