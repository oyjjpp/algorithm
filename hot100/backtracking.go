package hot100

import (
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
	res := make([][]int, 0)
	track := make([]int, 0)
	used := make([]bool, len(nums))

	var backtrack func(nums []int, track []int, used []bool)
	backtrack = func(nums, track []int, used []bool) {
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
			backtrack(nums, track, used)

			// 撤销选择
			track = track[:len(track)-1]
			used[i] = false
		}
	}
	backtrack(nums, track, used)
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
