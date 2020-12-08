package leetcode

import(
    "sort"
    "log"
)

// maxEnvelopes
// 俄罗斯套娃信封问题
func maxEnvelopes(envelopes [][]int) int {
    if len(envelopes)==0 || envelopes==nil {
        return 0    
    }
    
    // 对二维数组进行排序
    content := envelopesData{envelopes}
    sort.Sort(content)
    
    // 对高度数组寻找LIS
    height := make([]int, len(envelopes))
    for i:=0; i< len(envelopes); i++ {
        height[i] = content.data[i][1]
    }
    log.Println(height)
    
    return findMaxLengthOfLIS(height)
}

type envelopesData struct {
    data [][]int
} 

// Len
// 方法返回集合中的元素个数
func (p envelopesData) Len() int {
	return len(p.data)
}

// Less
// 方法报告索引i的元素是否比索引j的元素小
func (p envelopesData) Less(i, j int) bool {
    arr1 := p.data[i]
    arr2 := p.data[j]

    if arr1[0] < arr2[0]{
        return true
    } else if arr1[0] == arr2[0] {
        if arr1[1]>arr2[1]{
            return true
        }
    }
    return false
}

// Swap
// 方法交换索引i和j的两个元素
func (p envelopesData) Swap(i, j int) {
	p.data[i], p.data[j] = p.data[j], p.data[i]
}

// findMaxLengthOfLIS
// 最长递增子序列
// 使用DP Table
func findMaxLengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 初始化
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		// base case:dp数组全部初始化为1
		dp[i] = 1

		// 递增子序列，只要找到前面的那么比当前值小的子序列，然后接到最后，就会生成一个新的递增子序列
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		log.Println(dp[i])
	}
	// 排序 求最大的值
	sort.Ints(dp)
	log.Println(dp)
	return dp[len(dp)-1]
}

func max(a, b int) int{
    if a > b {
        return a
    }
    return b
}
