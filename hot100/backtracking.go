package hot100

import (
	"log"
	"sort"
	"strings"
)

// 回溯算法
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

// 46. 全排列
func permute(nums []int) [][]int {
	// 存储结果集
	res := make([][]int, 0)
	// 组合元素
	track := make([]int, 0)
	used := make([]bool, len(nums))

	var backtrack func(track []int, used []bool)
	backtrack = func(track []int, used []bool) {
		if len(track) == len(nums) {
			temp := make([]int, len(nums))
			copy(temp, track)
			res = append(res, temp)
			return
		}

		for i := range nums {
			// 排除不合法的选择
			if used[i] {
				continue
			}

			// 做选择
			track = append(track, nums[i])
			used[i] = true
			backtrack(track, used)

			// 撤销选择
			track = track[:len(track)-1]
			used[i] = false
		}
	}
	backtrack(track, used)
	return res
}

// 47. 全排列 II
func permuteUnique(nums []int) [][]int {
	// 存储结果集
	res := make([][]int, 0)
	// 组合元素
	track := make([]int, 0)
	used := make([]bool, len(nums))

	var backtrack func(track []int, used []bool)
	backtrack = func(track []int, used []bool) {
		if len(track) == len(nums) {
			temp := make([]int, len(nums))
			copy(temp, track)
			res = append(res, temp)
			return
		}

		for i := range nums {
			// 排除不合法的选择
			if used[i] {
				continue
			}

			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}

			// 做选择
			track = append(track, nums[i])
			used[i] = true
			backtrack(track, used)

			// 撤销选择
			track = track[:len(track)-1]
			used[i] = false
		}
	}
	sort.Ints(nums)
	backtrack(track, used)
	return res
}

// 排列（元素无重可复选）
func permuteRepeat(nums []int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)
	var backtrack func()
	backtrack = func() {
		if len(track) == len(nums) {
			temp := make([]int, len(nums))
			copy(temp, track)
			res = append(res, temp)
			return
		}

		for i := 0; i < len(nums); i++ {
			track = append(track, nums[i])
			backtrack()
			track = track[:len(track)-1]
		}
	}
	backtrack()
	return res
}

// 51. N 皇后 校验函数
func solveNQueens(n int) [][]string {
	// 校验是否为有效范围
	var isValid func(board []string, row, col int) bool
	isValid = func(board []string, row, col int) bool {
		// 检查列
		for i := 0; i < row; i++ {
			if board[i][col] == 'Q' {
				return false
			}
		}
		n := len(board)
		// 右上方
		for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
			if board[i][j] == 'Q' {
				return false
			}
		}
		for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
			if board[i][j] == 'Q' {
				return false
			}
		}
		return true
	}

	// 保存结果
	res := make([][]string, 0)
	board := make([]string, n)
	for i := 0; i < n; i++ {
		board[i] = strings.Repeat(".", n)
	}

	var backtrack func(board []string, row int)
	backtrack = func(board []string, row int) {
		if row == len(board) {
			newRow := make([]string, len(board))
			copy(newRow, board)
			res = append(res, newRow)
			return
		}

		n := len(board[row])
		for col := 0; col < n; col++ {
			// 校验是否符合要求
			if !isValid(board, row, col) {
				continue
			}

			newLine := []byte(board[row])
			newLine[col] = 'Q'
			board[row] = string(newLine)

			backtrack(board, row+1)

			newLine[col] = '.'
			board[row] = string(newLine)
		}
	}

	backtrack(board, 0)
	return res
}

// 78. 子集
func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)
	var backtrack func(start int)
	backtrack = func(start int) {
		temp := make([]int, len(track))
		copy(temp, track)
		res = append(res, temp)

		// 使用start 保证子集
		for i := start; i < len(nums); i++ {
			track = append(track, nums[i])
			log.Println(i, track)
			backtrack(i + 1)
			track = track[:len(track)-1]
		}
	}
	backtrack(0)
	return res
}

// 子集 II
func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)

	// 先排序，让相同的元素靠在一起
	sort.Ints(nums)

	var backtrack func(start int)
	backtrack = func(start int) {
		temp := make([]int, len(track))
		copy(temp, track)
		res = append(res, temp)

		// 使用start 保证子集
		for i := start; i < len(nums); i++ {

			// 值相同的 相邻节点 直接过滤
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			track = append(track, nums[i])
			log.Println(i, track)
			backtrack(i + 1)
			track = track[:len(track)-1]
		}
	}
	backtrack(0)
	return res
}

// 77. 组合
func combine(n int, k int) [][]int {
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, i+1)
	}
	res := make([][]int, 0)
	track := make([]int, 0)
	var backtrack func(start int)
	backtrack = func(start int) {
		if len(track) == k {
			temp := make([]int, k)
			copy(temp, track)
			res = append(res, temp)
			return
		}

		// 使用start 保证子集
		for i := start; i < n; i++ {
			track = append(track, nums[i])
			backtrack(i + 1)
			track = track[:len(track)-1]
		}
	}
	backtrack(0)
	return res
}

// 39. 组合总和
// 子集/组合（元素无重可复选
func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)
	trackSum := 0

	var backtrack func(start int)
	backtrack = func(start int) {
		if trackSum == target {
			temp := make([]int, len(track))
			copy(temp, track)
			res = append(res, temp)
		}
		if trackSum > target {
			return
		}

		for i := start; i < len(candidates); i++ {
			trackSum += candidates[i]
			track = append(track, candidates[i])

			backtrack(i)

			trackSum -= candidates[i]
			track = track[:len(track)-1]
		}
	}
	backtrack(0)
	return res
}
