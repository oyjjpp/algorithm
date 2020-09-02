package leetcode

func twoSum(nums []int, target int) []int {
    // 借助map实现
    data := map[int]int{}
    for k, v := range nums{
        temp := target-v
        if index,ok := data[temp];ok{
            return []int{index, k}
        }
        data[v]=k
    }
    return nil
}

