package leetcode

import (
	"log"
	"testing"
)

func TestMaxDistance(t *testing.T) {
	data := [][]int{
		{1, 0, 1},
		{0, 0, 0},
		{1, 0, 1},
	}

	rs := maxDistance(data)
	log.Println(rs)
}
