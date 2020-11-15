package leetcode

import (
	"log"
	"testing"
)

func TestCoinChange(t *testing.T) {
	rs := coinChangeV3([]int{2}, 1)
	log.Println(rs)
}

func TestSlice(t *testing.T) {
	data := []int{2}
	for _, value := range data {
		log.Println(value)
	}
}
