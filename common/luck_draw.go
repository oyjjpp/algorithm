package common

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"

	"github.com/spf13/cast"
)

type Gift struct {
	Name        string `json:"name"`
	Probability int    `json:"probability"`
}

// 构造容量为100的容器
func violence(gifts []Gift) int {
	length := 0
	data := ""
	for index, value := range gifts {
		length += value.Probability
		position := fmt.Sprintf("%d,", index)
		data += strings.Repeat(position, value.Probability)
	}

	// 获取[1,0) 随机数
	res, _ := rand.Int(rand.Reader, big.NewInt(100))
	index := res.Int64()

	arr := strings.Split(data, ",")
	giftIndex := cast.ToInt(arr[index])
	return giftIndex
}

// 离散算法
func dispersed(gifts []Gift) int {
	data := make([]int, 0)
	for index, value := range gifts {
		if index > 0 {
			data = append(data, value.Probability+data[index-1])
		} else {
			data = append(data, value.Probability)
		}
	}

	// 获取[1,0) 随机数
	result, _ := rand.Int(rand.Reader, big.NewInt(100))
	index := result.Int64()
	res := binarySearchV3(data, int(index))
	return res
}

// 二分法下界（重复元素 第一个元素）
func binarySearchV3(data []int, target int) int {
	left, right := 0, len(data)

	for left <= right {
		mid := left + (right-left)/2

		if data[mid] > target {
			right = mid - 1
		} else if data[mid] < target {
			left = mid + 1
		} else {
			if mid == 0 || data[mid-1] != target {
				return mid
			} else {
				right = mid - 1
			}
		}
	}

	return left
}

// 别名算法
func aliasMethod(gifts []Gift) int {
	// pdf := []float64{0.1, 0.2, 0.3, 0.4}
	// pdf := []float64{0.8, 0.1, 0.1}
	pdf := make([]float64, len(gifts))
	for index, value := range gifts {
		pdf[index] = float64(value.Probability) / 100
	}

	lenth := len(pdf)

	// 原始数据
	probInfo := make([]float64, lenth)

	// 别名补充数据
	alias := make([]int, lenth)

	small := []int{}
	large := []int{}

	// 构造拼装出每一列和都唯一的矩阵
	for i := 0; i < lenth; i++ {
		pdf[i] *= float64(lenth)
		if pdf[i] < 1.0 {
			small = append(small, i)
		} else {
			large = append(large, i)
		}
	}

	// 构造所有矩形概率值都等于1
	for len(small) != 0 && len(large) != 0 {
		s_index := small[0]
		small = small[1:]
		l_index := large[0]
		large = large[1:]

		probInfo[s_index] = pdf[s_index]
		alias[s_index] = l_index

		// 1.2 -= 1.0- 0.4
		pdf[l_index] -= 1.0 - pdf[s_index]
		if pdf[l_index] < 1.0 {
			small = append(small, l_index)
		} else {
			large = append(large, l_index)
		}
	}

	for len(small) > 0 {
		temp := small[0]
		small = small[1:]
		probInfo[temp] = 1.0
	}

	for len(large) > 0 {
		temp := large[0]
		large = large[1:]
		probInfo[temp] = 1.0
	}

	// 随机取出一个列
	result, _ := rand.Int(rand.Reader, big.NewInt(int64(lenth)))
	column := result.Int64()

	// 获取一个随机小数
	randData, _ := rand.Int(rand.Reader, big.NewInt(100))
	randomNumber := float64(randData.Int64()) / 100

	if randomNumber < probInfo[column] {
		return int(column)
	}
	return alias[column]
}
