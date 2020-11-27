package leetcode

// searchRange
// 在排序数组中查找元素的第一个和最后一个位置
func searchRange(nums []int, target int) []int {
    if len(nums) == 0 {
        return []int{-1, -1}
    }
    rs := []int{}
    left, right := 0, len(nums)
    
    for left < right {
        mid := (left + right)/2
        
        if nums[mid] == target {
            right = mid
        } else if nums[mid] < target {
            left = mid + 1
        } else if nums[mid] > target {
            right = mid
        }
    }
    // 左侧边际
    if left == len(nums) || nums[left] != target {
        return []int{-1, -1}
    }
    rs = append(rs, left)
    for i:=left; i<len(nums)-1; i++ {
        if nums[i] == target{
            right = i
        }else {
            break
        }
    }
    rs = append(rs, right)
    return rs
}

// leftSearch
// 寻找左侧边际的二分查找
func leftSearch(nums []int, target int) int {
    if len(nums) == 0 {
        return -1
    }
    left, right := 0, len(nums)-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] < target {
            left = mid + 1
        } else if nums[mid] > target {
            right = mid - 1
        }else if nums[mid] == target {
            right = mid - 1
        }
    }
    
    if left == len(nums) || nums[left] != target {
        return -1
    }
    return left
}
