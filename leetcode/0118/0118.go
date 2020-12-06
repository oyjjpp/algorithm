package leetcode 

// 杨辉三角
// 1、col=0  元素为1
// 2、col=row 元素为1 并不在为当前行余下元素赋值
func generate(numRows int) [][]int {
	if numRows < 0 {
		return [][]int{}
	}

	//声明一个二维切片
	res := make([][]int, numRows)
	for index := 0; index < numRows; index++ {
		res[index] = make([]int, index+1)
	}

	for i := 0; i < numRows; i++ {
		for j := 0; j < numRows; j++ {
			if j == i {
                // 对角线位置
				res[i][j] = 1
				break
			} else if j == 0 {
                // 第一列
				res[i][j] = 1
			} else if i > 0 && j > 0 {
				res[i][j] = res[i-1][j] + res[i-1][j-1]
			}
		}
	}
	return res
}
