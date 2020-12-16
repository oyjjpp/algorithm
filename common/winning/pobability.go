package winning

import (
    math/rand"
)

// 普通概率模型
func roll(probability float) {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Float32() <= probability;
}

