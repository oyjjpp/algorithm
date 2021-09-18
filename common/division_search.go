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

