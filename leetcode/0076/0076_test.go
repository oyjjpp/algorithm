package leetcode

import (
	"encoding/json"
	"testing"
    "reflect"
)

func TestMapData(t *testing.T) {
	data := "avavde"
    t.Log(reflect.ValueOf(data[0]).Kind())
	window := map[byte]int{}
	for i := 0; i < len(data); i++ {
		value := data[i]
        t.Log(reflect.ValueOf(value).Kind())
		window[value]++
	}
	rs, err := json.Marshal(window)
	if err == nil {
		t.Log(string(rs))
	}
}

func TestMinWindow(t *testing.T) {
	// 即包含一个长的右包含一个短的满足条件的子字符串
    rs := minWindow("ADOBECODEBANC", "ABC")
    t.Log(rs)
    
    // 全部匹配
    rs = minWindow("a", "a")
    t.Log(rs)
    
    // 只包含一个：当前于全匹配
    rs = minWindow("ADOBEC", "ABC")
    t.Log(rs)
    
    // 前面包含无效字符：通过窗口缩小去掉
    // 如何剔除前面无效字符的
    // 问题：
    // 1，如何剔除的非满足元素？
    // 2，索引是如何更新的？
    rs = minWindow("FADOBEC", "ABC")
	t.Log(rs)
    
    
    // 后面包含无效字符：right到达s的尽头
    rs = minWindow("ADOBECF", "ABC")
	t.Log(rs)
}
