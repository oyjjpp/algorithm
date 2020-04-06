package common

import (
	"testing"
)

func TestSort(t *testing.T) {
	//data := []int{9, 3, 5, 4, 9, 1, 2, 7, 8, 1, 3, 6, 5, 3, 4, 0, 10, 9, 7, 9}
	data := []int{79, 73, 75, 74, 79, 71, 72, 77, 88, 81, 83, 86, 85, 83, 84, 80, 80, 89, 87, 89}
	//data := []int{10000, 0}
	rs := CountSort(data)
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

func TestString(t *testing.T) {
	rs := fractionToDecimal(23, 6)
	t.Log(rs)
}

func TestArray(t *testing.T) {
	cont := Constructor([]int{1, 2, 3})
	t.Log(cont.ShuffleV2())
	t.Log(cont.ShuffleV2())
	t.Log(cont.ShuffleV2())
}

func TestList(t *testing.T) {
	//[1,0,3,4,0,1]
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 1,
						},
					},
				},
			},
		},
	}
	res := sortList(head)
	PrintList(res)
}
