package main

import "fmt"

func main() {
	//{0} 0
	//{1} 1

	//{1,2,3,0,0,0} 3
	//{2,5,6} 3

	//{2,0} 1
	//{1} 1

	//{4, 5, 6, 0, 0, 0} 3
	//{1, 2, 3} 3

	// 4,0,0,0,0,0
	//1,2,3,5,6
	nums1 := []int{4, 0, 0, 0, 0, 0}
	nums2 := []int{1, 2, 3, 5, 6}
	merge(nums1, 1, nums2, 5)
	fmt.Println(nums1)
}

// merge
// 合并两个有序数组
func merge(nums1 []int, m int, nums2 []int, n int) {
	indexm, indexn := m-1, n-1
	for i := m + n - 1; i >= 0; i-- {
		// 无nums1元素 将nums2赋值nums1
		if indexm < 0 {
			nums1[i] = nums2[indexn]
			indexn--
			continue
		}
		// 无nums2元素，将nums1赋值nums1
		if indexn < 0 {
			nums1[i] = nums1[indexm]
			indexm--
			continue
		}

		// 都存在时 哪个大哪个先赋值nums1
		if nums1[indexm] >= nums2[indexn] {
			nums1[i] = nums1[indexm]
			indexm--
			continue
		}
		if nums1[indexm] < nums2[indexn] {
			nums1[i] = nums2[indexn]
			indexn--
			continue
		}
	}
}
