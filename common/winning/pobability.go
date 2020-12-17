package winning

import (
    "time"
    "math/rand"
)

// roll
// @param probability 概率
// 普通概率模型
// 策划希望平均每10次抽奖能中一次，则每次抽奖中奖率为10%
func roll(probability float32) bool {
    // 注册随机种子
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Float32() <= probability;
}

var count int
// roolV1
// 固定概率
// @param probability 概率
// @param threshold
func roolV1(probability float32, threshold int) bool {
    // 累计次数
    count++
    // 达到累计次数
    if count % threshold == 0 {
        return true
    } else {
        return roll(probability)   
    }
}

// roolV2
// 计数器模型
// @param probability 概率
// @param threshold
func roolV2(probability float32, threshold int) bool {
    // 累计次数
    count++
    result := false 
    // 达到累计次数
    if count % threshold == 0 {
        result = true
    } else {
        result = roll(probability)   
    }
    // 恢复基数
    if result {
        count = 0
    }
    return result
}
