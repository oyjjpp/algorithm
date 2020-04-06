// 数组相关算法
package common

import (
	"fmt"
	"math"
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
// 334、递增的三元子序列
// 238、除自身以外数组的乘积

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

	// 根据k进行两次翻转
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
		// 选取一个元素与其他进行比较
		for j := length - 1; j > i; j-- {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

// containsDuplicateV2
// 借助hash唯一性解决,通过空间置换时间
func containsDuplicateV2(nums []int) bool {
	length := len(nums)
	// 创建一个map key存储当前元素值
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
// 思路
// 1、使用两个计数器，一个用来递增校验，一个用于递减默认填充零
// 2、临界值：两个计数器碰撞到一起
func moveZeroes(nums []int) {
	l := len(nums) //递减非零计数器
	i := 0         //递增非零计数器
	for {
		// 两个计数器碰撞到一起 则说明已经全部移动完成
		if i >= l {
			break
		}
		// [4,1,2,0,0,1]
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
// 借助hashMap实现
func intersect(nums1 []int, nums2 []int) []int {
	// 将长的数组放入到hashMap中
	if len(nums1) < len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	hashMap := map[int]int{}

	// 将第一个元素存储到hashmap中
	for _, value := range nums1 {
		// 已经存在
		if _, ok := hashMap[value]; ok {
			hashMap[value]++
		} else {
			hashMap[value] = 1
		}
	}
	// 用于存储结果
	res := []int{}

	// 通过另一个数组校验
	for _, value := range nums2 {
		if _, ok := hashMap[value]; ok {
			res = append(res, value)
			hashMap[value]--
			// hashMap中数据已经减少到0 则需要将hashMap的key unset掉
			if hashMap[value] == 0 {
				delete(hashMap, value)
			}
		}
	}
	return res
}

// increasingTriplet
// 递增的三元子序列
// 给定一个未排序的数组，判断这个数组中是否存在长度为3的递增子序列。
func increasingTriplet(nums []int) bool {
	// 记录长度
	length := 0
	// 记录当前元素
	curEle := 0
	// [5,1,5,5,2,5,4]
	for index, value := range nums {
		if index == 0 {
			length++
			curEle = value
		} else if length >= 3 {
			return true
		} else if curEle < value {
			length++
			curEle = value
			if length >= 3 {
				return true
			}
		} else {
			length = 1
			curEle = value
		}
	}
	return false
}

// increasingTriplet
// 递增的三元子序列【不需要连续】
// 思路
// 1、a 始终记录最小元素，b 为某个子序列里第二大的数;
// 2、接下来不断更新 a，同时保持 b 尽可能的小;
// 3、如果下一个元素比b大，说明找到了三元组;
func increasingTripletV2(nums []int) bool {
	one := math.MaxInt32
	two := math.MaxInt32
	// [2,4,-2,-3]
	for _, value := range nums {
		if value <= one {
			one = value
		} else if value <= two {
			two = value
		} else {
			return true
		}
	}
	return false
}

// 除自身以外数组的乘积
// 给你一个长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，
// 其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。
// 思路
// 1、先计算当前元素左侧的乘积
// 2、再计算当前元素右侧的乘积，两者相乘
func productExceptSelf(nums []int) []int {
	length := len(nums)
	res := make([]int, length)
	right, left := 1, 1
	// 先计算左侧乘积
	for i := 0; i < length; i++ {
		res[i] = left
		left *= nums[i]
	}
	// 左侧与右侧相乘
	for i := length - 1; i >= 0; i-- {
		res[i] *= right
		right *= nums[i]
	}
	return res
}
