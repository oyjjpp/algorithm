// 常见排序算法
package common

// 二路归并

//@link 排序算法 https://juejin.im/post/5a08cc646fb9a045030f9174

/**
 * 插入排序 {5, 2, 0, 1, 3, 1, 4}
 * 从一组元素中取一个元素为有序元素组，然后在剩下的元素中每次取一个元素向有序的元素组插
 * 时间复杂度O(n^2)
 * 空间复杂度O(1)
 * 稳定性：稳定
 */
func InsertSort(item []int) []int {
	for i := 1; i < len(item); i++ {
		// 选取一个值 向有序集合中插入
		temp := item[i]
		j := i - 1

		for ; j >= 0 && item[j] > temp; j-- {
			item[j+1] = item[j]
		}
		item[j+1] = temp
	}
	return item
}

/**
 * @desc 希尔排序
 * 时间复杂度O(n^1.3)
 * 空间复杂度O(1)
 * 稳定性：不稳定
 */
func ShellSort(item []int) []int {
	var temp int
	var j int

	n := len(item)
	for d := n / 2; d > 0; d = d / 2 {
		for x := 0; x < d; x++ {
			for i := x + d; i < n; i = i + d {
				temp = item[i]
				j = i - d
				for ; j >= 0 && item[j] > temp; j = j - d {
					item[j+d] = item[j]
				}
				item[j+d] = temp
			}
		}
	}
	return item
}

/**
 * 简单选择排序
 * 循环查找“最小”元素放在首位
 * 时间复杂度O(n^2)
 * 空间复杂度O(1)
 * 稳定性 : 不稳定
 */
func SelectSort(item []int) []int {
	var j int
	var temp int
	var position int

	for i, n := 0, len(item); i < n; i++ {
		j = i + 1
		temp = item[i]
		position = i
		for ; j < n; j++ {
			if item[j] < temp {
				temp = item[j]
				position = j
			}
		}
		item[position] = item[i]
		item[i] = temp
	}

	return item
}

/**
 * @desc 创建最大堆
 * @param slice item 元素组
 * @param heapSize int 需要创建最大堆的大小
 * @param index int 当前需要创建最大堆的位置
 */
func maxHeapify(item []int, heapSize, index int) {
	left := index*2 + 1
	right := left + 1

	largest := index

	if left < heapSize && item[index] < item[left] {
		largest = left
	}

	if right < heapSize && item[largest] < item[right] {
		largest = right
	}

	if largest != index {
		temp := item[index]
		item[index] = item[largest]
		item[largest] = temp
		maxHeapify(item, heapSize, largest)
	}
}

/**
 * @desc 堆排序
 * @param 时间复杂度 O(nlogn)
 */
func HeadSort(item []int) []int {
	n := len(item)
	startIndex := (n - 1 - 1) / 2

	for i := startIndex; i >= 0; i-- {
		maxHeapify(item, n, i)
	}

	var temp int
	for i := n - 1; i > 0; i-- {
		temp = item[0]
		item[0] = item[i]
		item[i] = temp
		maxHeapify(item, i, 0)
	}

	return item
}

// 冒泡排序
// BubbleSort
// 原理
// 两两比较相邻记录的排序码，若发生逆序，则交换
// 时间复杂度O(n^2) 空间复杂度O(1)
func BubbleSort(item []int) []int {
	n := len(item)

	for i := 0; i < n-1; i++ {
		for j := n - 1 - 1; j >= i; j-- {
			if item[j+1] < item[j] {
				temp := item[j]
				item[j] = item[j+1]
				item[j+1] = temp
			}
		}
	}
	return item
}

/**
 * @desc 快速排序
 * 时间复杂度O(nlogn)
 * 算法：选择一个基数，一般我们选择第一个数，然后把大于该数的放右边，小于该数的放左边，然后分别对左右两边用同样的方法处理，直到排序结束。
 */
func QuikcSort(item []int) []int {
	quickSort(item, 0, len(item)-1)
	return item
}

//交换
func swap(item []int, i, j int) {
	temp := item[i]
	item[i] = item[j]
	item[j] = temp
}

//@desc 快速排序
//@param item 待排序的数组
//@param start 开始位置
//@param end 结束位置
func quickSort(item []int, start, end int) {
	if start < end {
		//第一个元素作为基数
		pivot := item[start]
		left := start
		right := end

		for left != right {
			//最右边的元素大于基数
			for item[right] >= pivot && left < right {
				right--
			}

			for item[left] <= pivot && left < right {
				left++
			}
			swap(item, left, right)
		}

		item[start] = item[left]
		item[left] = pivot
		quickSort(item, start, left-1)
		quickSort(item, left+1, end)
	}
}

// QuertSortV2
// 快速排序 递归 取一个值作为基准,大于基数的放在左边，小于基于的放在右边
func QuertSortV2(data []int) []int {
	if len(data) < 1 {
		return data
	}
	// 获取第一个为基数
	temp := data[0]

	left := []int{}
	right := []int{}
	for i := 1; i < len(data); i++ {
		if data[i] > temp {
			right = append(right, data[i])
		} else {
			left = append(left, data[i])
		}
	}
	left = QuertSortV2(left)
	right = QuertSortV2(right)

	return append(append(left, temp), right...)
}

/**
 * @desc 归并排序
 * 时间复杂度 O(nlog2n)
 * 空间复杂度 O(n) + O(log2n)
 * 稳定性：稳定
 */
func MergeSort(item []int) []int {
	mergeSort(item, 0, len(item)-1)
	return item
}

// @param item 排序数组
// @param 开始索引位置
// @param 结束索引位置
func mergeSort(item []int, left, right int) {
	if left < right {
		center := (left + right) / 2
		mergeSort(item, left, center)
		mergeSort(item, center+1, right)
		merge(item, left, center+1, right)
	}
}

// @desc 合并两个数组
// @link https://juejin.im/post/5ab4c7566fb9a028cb2d9126
func merge(item []int, left, center, right int) {
	// 左侧数组大小
	leftData := make([]int, center-left)
	// 右侧数组大小
	rightData := make([]int, right-center+1)

	// 向两个数组中填充数据
	for i := left; i < center; i++ {
		leftData[i-left] = item[i]
	}

	for i := center; i <= right; i++ {
		rightData[i-center] = item[i]
	}

	// 用于遍历两个数组
	i, j := 0, 0
	// 数组中的第一个元素
	index := left
	// 循环对比合并两个数组
	for i < len(leftData) && j < len(rightData) {
		if leftData[i] < rightData[j] {
			item[index] = leftData[i]
			i++
		} else {
			item[index] = rightData[j]
			j++
		}
		// 增加后索引增加1
		index++
	}

	// 将数据中剩余的元素继续插入
	for i < len(leftData) {
		item[index] = leftData[i]
		i++
		index++
	}
	for j < len(rightData) {
		item[index] = rightData[j]
		j++
		index++
	}
}

func RadixSort(item []int) []int {
	max := item[0]
	itemLen := len(item)

	for i := 1; i < itemLen; i++ {
		if item[i] > max {
			max = item[i]
		}
	}

	time := 0 //数组最大值位数
	for max > 0 {
		max = max / 10
		time++
	}

	return item
}
