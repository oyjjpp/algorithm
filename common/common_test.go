package common

import "testing"

func TestSort(t *testing.T) {
	//data := []int{335, 2, 0, 16, 3, 2326, 34, 23, 1, 4}
	data := []int{10000, 0}
	rs := RadixSort(data)
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
