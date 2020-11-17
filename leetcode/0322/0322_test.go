package leetcode

import (
	"log"
	"testing"
)

func TestCoinChange(t *testing.T) {
	// [3] 2
	// [2] 1
	// [1,2,5] 11
	// [2,5,10,1] 27
	rs := coinChangeV3([]int{1, 2, 5}, 11)
	log.Println(rs)
}

func TestSlice(t *testing.T) {
	data := []int{2}
	for _, value := range data {
		log.Println(value)
	}
}
