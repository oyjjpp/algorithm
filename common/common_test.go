package common

import "testing"

func TestBubbleSort(t *testing.T) {
	item := []int{12, 4, 1, 523, 6}
	rs := BubbleSort(item)
	t.Log(rs)
}

func TestSort(t *testing.T) {
	data := []int{5, 2, 0, 1, 3, 1, 4}
	rs := QuertSortV2(data)
	t.Log(rs)
}

func InsertSortV2(data []int) []int {
	var temp int
	for i := 1; i < len(data); i++ {
		// 从未排序中选择其中一个
		temp = data[i]
		j := i - 1
		for ; j >= 0 && data[j] > temp; j-- {
			data[j+1] = data[j]
		}
		data[j+1] = temp
	}
	return data
}
