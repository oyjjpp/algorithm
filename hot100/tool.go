package hot100

import "log"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// min
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Common() {
	rs := 1 / 2
	log.Println(rs)
}
