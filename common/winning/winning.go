package winning

import (
    "encoding/json"
    "math/rand"
	"strconv"
	"time"
)

// 方法一
// 扩展集合，使每一项出现的次数与其权重正相关
// 优点：选取的时间复杂度为O（1）,算法简单。
// 缺点：空间占用极大，如果权重数字位数较大，例如{A:49.1 B：50.9}的时候，就会产生巨大的空间浪费。
func probability(config string) string {
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
// 产生一个随机整数 0-5
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
func probabilityV2() string{
    
}
