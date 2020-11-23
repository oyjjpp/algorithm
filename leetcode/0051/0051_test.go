package leetcode

import (
    "testing"
    "log"
    "fmt"
    "encoding/json"
)

func TestSolveNQueens(t *testing.T){
    rs := solveNQueens(4)
    res, _ := json.Marshal(rs)
    log.Println(string(res))
}

func TestTwoSlice(t *testing.T) {
    n := 8
	track := make([][]string, n, n)

	// 初始化路径track "."，选择后改为"Q"
	for i := 0; i < n; i++ {
		tmp := make([]string, n, n)
		for j := 0; j < n; j++ {
            tmp[j] = "."
			// tmp = append(tmp, ".")
		}
        track[i] = tmp
		//track = append(track, tmp)
	}
    rs, _ := json.Marshal(track)
    log.Println(string(rs))
    fmt.Println()
    for _, value := range track {
        for _, item := range value {
            fmt.Printf(item)
            fmt.Printf("-")
        }
        fmt.Println()
    }
}
