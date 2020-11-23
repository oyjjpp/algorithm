package leetcode

import (
	"log"
	"testing"
)

func TestTwoSlice(t *testing.T) {
	data := [][]int{
		{1, 2, 3},
		{3, 2, 1},
	}
	log.Println(data, len(data))
}

func TestPermute(t *testing.T) {
	// 5,4,2,6
	// 1,2,3
	data := []int{5, 4, 2, 6}
	res := permuteV2(data)
	log.Println(res)
}

/*
result = []
def backtrack(路径, 选择列表):
    if 满足结束条件:
        result.add(路径)
        return

    for 选择 in 选择列表:
        做选择
        backtrack(路径, 选择列表)
		撤销选择
*/

// permuteV2
// 全排列
// 回溯问题
// 产生问题
func permuteV2(nums []int) [][]int {
	result := [][]int{}

	// 是否存在
	visited := map[int]bool{}

	var backtrack func(path []int)
	backtrack = func(path []int) {
		// 满足结束条件
		if len(path) == len(nums) {
			// TODO 直接赋值 与 copy方式赋值结果不一致
			// 原因：切片赋值会直接拷贝原数组的地址，如果原数组发生了变化，赋值后的元素也会发生变化
			// 使用copy因为需要新创建一个切片，使用了新创建的地址，所以如此问题
			// temp := path

			// 通过copy方式赋值
			temp := make([]int, len(path))
			copy(temp, path)

			result = append(result, temp)
			log.Println("temp", temp, &path, result)
			return
		}

		for _, index := range nums {
			// 做选择
			if bo, ok := visited[index]; ok && bo {
				continue
			}
			visited[index] = true
			path = append(path, index)
			backtrack(path)
			log.Println(index, path)
			// 撤销选择
			visited[index] = false
			path = path[:len(path)-1]
			if len(path) == 0 {
				log.Println("一轮结束")
			}
		}

	}
	backtrack([]int{})
	return result
}

func TestCopySlice(t *testing.T) {
	s := []int{1, 2, 3}
	log.Println(s)
	copy(s, []int{4, 5, 6, 7, 8, 9})
	log.Println(s)
}

func TestSlice(t *testing.T) {
	path := []int{1, 2, 3}
	temp1 := path
	copy1 := make([]int, len(path))
	copy(copy1, path)
	log.Println(temp1, copy1)

	path[2] = 0
	log.Println(temp1, copy1)
}
