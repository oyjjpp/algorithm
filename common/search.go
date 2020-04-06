package common

// 162、寻找峰值
// 287、寻找重复数
// 315、计算右侧小于当前元素的个数

// findPeakElement
// 寻找峰值
func findPeakElement(nums []int) int {
	left := 0
	right := len(nums) - 1
	// [1,2,3,1]
	for left < right {
		mid := (left + right) / 2
		// 左边高，说明左边有峰值，可能mid就是
		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			// 右边高，说明在mid右边有峰值，所以mid一定不是
			left = mid + 1
		}
	}
	return left
}

// findDuplicate
// 寻找重复数
// 给定一个包含 n + 1 个整数的数组nums，其数字都在1到n之间（包括1和n），可知至少存在一个重复的整数。假设只有一个重复的整数，找出这个重复的数。
// 解法跟链表寻找环一样，都是快慢指针，来判断是有环的，然后寻找到他们的相遇点，然后让low从0开始，fast从相遇点开始
func findDuplicate(nums []int) int {
	low, fast := nums[0], nums[nums[0]] // 首先制定固有的速率，low是nums[low]的速度，fast是二倍于它的速度就是 nums[nums[0]]
	for low != fast {
		low = nums[low] // 按照不同的速度（存在倍数关系），如果这两者能相遇，就只有一个原因就是因为存在环。
		fast = nums[nums[fast]]
	}
	// 但是他们相遇的地方不一定是入口处，有可能是环的内部的某个元素，这个元素叫做相遇点。
	low = 0 // 然后这一步我门要寻找这个环的入口，将low设置为0就是从头开始走，然后fast在相遇点。这个时候把他们的速度调成一样。
	// 即： low = nums[low] fast = nums[fast] 这样速度就一样了，（都是一次，之前速度不同的时候fast直接两次）
	for low != fast {
		low = nums[low]
		fast = nums[fast]
	}
	return low
}
