// 数组相关算法
package common

import (
	"fmt"
	"math/rand"
)

// leetcode
// 152、乘积最大子数组
// 169、多数元素【求众数】
// 189、旋转数组
// 217、存在重复元素
// 283、移动零
// 384、打乱数组
// 350、两个数组的交集||

// maxProduct
// 乘积最大子数组
// 给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字）
// 思路
// 1、当一个数组中没有0存在，则分为两种情况
// 2、负数为偶数个，则整个数组的各个值相乘为最大值
// 3、负数为奇数个，则从左边开始，乘到最后一个负数停止有一个“最大值”,从右边开始也有一个“最大值”
// 进行比较即可以得出最大值
func maxProduct(nums []int) int {
	number := 1
	max := nums[0]

	// 左起开始
	for _, value := range nums {
		// 一直保存所有数乘机与最大值做对比
		number = number * value
		if max < number {
			max = number
		}
		// 当出现0 则从新开始
		if number == 0 {
			number = 1
		}
	}

	// 右起开始
	number = 1
	for i := len(nums) - 1; i >= 0; i-- {
		number = number * nums[i]
		if max < number {
			max = number
		}
		if number == 0 {
			number = 1
		}
		fmt.Println("max", max)
	}
	return max
}

// majorityElement
// 多数元素【求众数】
// 给定一个大小为 n 的数组，找到其中的多数元素;多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素
// 思路
// 摩尔投票法
// 1、抗阶段：分属两个候选人的票数进行两两对抗抵消
// 2、计数阶段：计算对抗结果中最后留下的候选人票数是否有效
func majorityElement(nums []int) int {
	count := 0
	major := 0
	for _, value := range nums {
		if count == 0 {
			// 完全消除掉时，则重新复制
			major = value
			count++
		} else if major == value {
			// 相等则加一
			count++
		} else {
			// 不相等则抵消
			count--
		}
	}
	return major
}

// rotate
// 旋转数组
// 思路
// 1、通过所有元素旋转
// 2、旋转指定长度
func rotate(nums []int, k int) {
	// 翻转
	reverse := func(nums []int, start, end int) {
		for start < end {
			temp := nums[start]
			nums[start] = nums[end]
			start++
			nums[end] = temp
			end--
		}
	}
	// [1,2,3,4,5,6,7] 3
	// 数组长度
	length := len(nums)
	k = k % length
	// 整理翻转
	reverse(nums, 0, length-1)

	// 根据进行两次翻转
	reverse(nums, 0, k-1)
	reverse(nums, k, length-1)
}

// 旋转数组
// 双循环
func rotateV2(nums []int, k int) {
	length := len(nums)
	k = k % length

	for i := 0; i < k; i++ {
		// 最后一个元素,放在第一个位置
		temp := nums[length-1]
		// 其他元素顺序先后移动一个位置
		for j := length - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = temp
	}
}

// containsDuplicate
// 存在重复元素
// 给定一个整数数组，判断是否存在重复元素
// 暴力破解
func containsDuplicate(nums []int) bool {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := length - 1; j > i; j-- {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

// containsDuplicateV2
// 借助hash唯一性解决
func containsDuplicateV2(nums []int) bool {
	length := len(nums)
	data := map[int]int{}
	for i := 0; i < length; i++ {
		if _, ok := data[nums[i]]; ok {
			return true
		} else {
			data[nums[i]] = i
		}
	}
	return false
}

// moveZeroes
// 移动零
func moveZeroes(nums []int) {
	l := len(nums) //递减非零计数器
	i := 0         //递增非零计数器
	for {
		if i >= l {
			break
		}
		if nums[i] == 0 {
			// 遇到一个0，减少1，最后与i比较作为结束条件
			l = l - 1
			// 把0前后两部分合并
			nums = append(nums[0:i], nums[i+1:]...)
			// 在末尾补回0
			nums = append(nums, 0)
		} else {
			i = i + 1 //非0计数器自增
		}
	}
	return
}

type Solution struct {
	// 为了保存初始值
	nums []int
}

// 创建一个Solution
func Constructor(nums []int) Solution {
	return Solution{nums}
}

// 对数组进行重置
func (this *Solution) Reset() []int {
	return this.nums
}

// 随机打散数组元素
func (this *Solution) Shuffle() []int {
	nums := make([]int, len(this.nums))
	// 拷贝一份数组，不影响原有数组
	copy(nums, this.nums)
	// 使用官方库处理
	rand.Shuffle(len(nums), func(i int, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
	return nums
}

// 通过rand.Intn(len(nums)) 引入随机数
func (this *Solution) ShuffleV2() []int {
	nums := make([]int, len(this.nums))
	copy(nums, this.nums)
	for i := range nums {
		j := rand.Intn(len(nums))
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}

// intersect
// 两个数组的交集 II
func intersect(nums1 []int, nums2 []int) []int {
	return []int{}
}
