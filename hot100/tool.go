package hot100

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
