package common

import "testing"

func TestBubbleSort(t *testing.T) {
	item := []int{12, 4, 1, 523, 6}
	rs := BubbleSort(item)
	t.Log(rs)
}
