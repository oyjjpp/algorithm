package leetcode

import (
	"encoding/json"
	"testing"
)

func TestBuildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	data := buildTree(preorder, inorder)
	rs, err := json.Marshal(*data)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(string(rs))
}
