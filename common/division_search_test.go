package common

import (
	"testing"
)

func TestBinarySearch(t *testing.T){
	data := []int{1,3,6,8,10}
	rs := binarySearch(data,6)
	t.Log(rs)
}
