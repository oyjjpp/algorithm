package common

import (
	"testing"
)

func TestBinarySearch(t *testing.T){
	data := []int{1,3,6,6,6,8,10}
	rs := binarySearchV1(data,6)
	t.Log(rs)
}
