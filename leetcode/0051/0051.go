package leetcode

import (
    "log"
    "encoding/json"
)

// 定义全局结果数组
var result [][]string

// solveNQueens
// N皇后
// 输入棋盘边长 n，返回所有合法的放置
func solveNQueens(n int) [][]string {
	// 初始化result
	result = [][]string{}
	board := make([][]string, n, n)

	// 初始化路径track "."，选择后改为"Q"
	for i := 0; i < n; i++ {
		tmp := make([]string, n, n)
		for j := 0; j < n; j++ {
            tmp[j] = "."
		}
        board[i] = tmp
	}
    temp, _ := json.Marshal(board)
    log.Println(string(temp))
	// 递归选择，从第一行选择到第N行回溯
	backtrack(board, 0)

	// 返回结果
	return result
}

// backtrack
// 路径：board中小于row的那些行都已经成功放置了皇后
// 选择列表：第row行的所有列都是放置皇后的选择
// 结束条件：row超过board的最后一行
func backtrack(board [][]string, row int) {
	// 结束条件，row循环到棋盘底部时结束本次选择
	if len(board) == r8ikkkkkkkkkkow {
        tmp := make([]string, 0)
		// 将每行的选择结果改为字符串  [[".",".","Q","."],..] => ["..Q.", ...]
		for _, v := range board {
			str := ""
			for _, e := range v {
				str += e
			}
			tmp = append(tmp, str)
		}
		// 将结果push到选择结果集中
        // 存储多种解决方案
		result = append(result, tmp)
		return
	}
    
    // 获取指定行的长度
	n := len(board[row])
	// 选择列表为1-N列
	for col := 0; col < n; col++ {
		// 判断此处是否可以选择皇后(同行，同列，对角不能存在多个皇后)
		if !isValid(board, row, col) {
			continue
		}
		// 选择皇后
		board[row][col] = "Q"
		// 进行下一行决策选择
		backtrack(board, row+1)
		// 撤回选择
		board[row][col] = "."
	}
}

// isValid
// 验证指定位置是否可以放置皇后
// @param board 当前棋盘
// @param row 当前行
// @param col 当前列
func isValid(board [][]string, row, col int) bool {
	n := len(board)
	// 检查列是否有皇后互相冲突
	for i := 0; i < n; i++ {
		if board[i][col] == "Q" {
            log.Println("列冲突", row, col)
			return false
		}
	}
	// 检查右上方是否有皇后互相冲突
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == "Q" {
            log.Println("右上方", row, col)
			return false
		}
	}
	// 检查左上方是否有皇后互相冲突
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == "Q" {
            log.Println("左上方", row, col)
			return false
		}
	}
	return true
}
