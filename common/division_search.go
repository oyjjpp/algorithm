package common

func binarySearch(data []int, target  int) int {
	left,right := 0 ,len(data)

	for left <= right {
		mid := left + (right-left)/2

		if data[mid] == target{
			return mid
		} else if data[mid] < target{
			left = mid+1
		}else{
			right = mid-1
		}
	}

	return -1;
}

// 二分法下界（重复元素 第一个元素）
func binarySearchV1(data []int, target  int) int {
	left,right := 0 ,len(data)

	for left <= right {
		mid := left + (right-left)/2

		if data[mid]>target {
			right = mid-1
		}else if data[mid]<target {
			left = mid+1
		}else{
			if mid ==0 || data[mid-1]!=target{
				return mid
			}else{
				right = mid-1
			}
		}
	}

	return -1
}

// 二分法上界（重复元素 最后一个元素）
func binarySearchV2(data []int, target  int) int {
	left,right := 0 ,len(data)

	length := right

	for left <= right {
		mid := left + (right-left)/2

		if data[mid]>target {
			right = mid-1
		}else if data[mid]<target {
			left = mid+1
		}else{
			if mid ==length-1 || data[mid+1]!=target{
				return mid
			}else{
				left = mid+1
			}
		}
	}

	return -1
}
