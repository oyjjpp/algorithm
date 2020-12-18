package winning

import (
    "encoding/json"
    "math/rand"
	"strconv"
	"time"
)

// 方法一
// 扩展集合，使每一项出现的次数与其权重正相关
// 根据元素权重值 "A"*2 ..等，把每个元素取权重个元素放到一个数组中，然后获取一个[0,n-1]随机值,到数组中获取 
// 优点：选取的时间复杂度为O（1）,算法简单。
// 缺点：空间占用极大，如果权重数字位数较大，例如{A:49.1 B：50.9}的时候，就会产生巨大的空间浪费。
func randomPobabilityV1(config string) string {
    // {"A":5, "B":2, "C":2, "D":1}
    var data map[string]int
    if err := json.Unmarshal([]byte(config), &data);err != nil{
        // 异常
        return ""
    }
    
    max := 0
    container := []string{}
    for ele, nums := range data {
        max = max + nums
        for i:=0 ;i<nums ;i++{
            container = append(container, ele)
        }
    }
    rand := Rand(0, max)
    return container[rand]
}

// Rand
// 产生一个随机整数[min, max)
func Rand(min, max int) int {
	if min > max {
		panic("min: min cannot be greater than max")
	}
	if int31 := 1<<31 - 1; max > int31 {
		panic("max: max can not be greater than " + strconv.Itoa(int31))
	}
	if min == max {
		return min
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min) + min
}

// 方法二
// 优化思路：减少内存的消耗
// 使用权重之和计算处一个随机数
// 
func randomPobabilityV2(config string) string{
    var data map[string]int
    if err := json.Unmarshal([]byte(config), &data);err != nil{
        // 异常
        return ""
    }
    // 计算当前权重总和，求出一个随机数
    sum := 0
    for _, widght := range data {
        sum = sum + widght
    }
    rand := Rand(1, sum+1)
    
    curWigth := 0 
    for key, widght := range data {
        curWigth = curWigth + widght
        if curWigth>=rand {
            return key
        }
    }
    return ""
}
