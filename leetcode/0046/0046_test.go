package leetcode

import (
	"log"
	"testing"
)

func TestTwoSlice(t *testing.T) {
	data := [][]int{
		{1, 2, 3},
		{3, 2, 1},
	}
	log.Println(data, len(data))
}

func TestPermute(t *testing.T) {
	data := []int{1, 2, 3}
	res := permute(data)
	log.Println(res)
}
