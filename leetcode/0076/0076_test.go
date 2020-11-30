package leetcode

import (
	"encoding/json"
	"testing"
)

func TestMapData(t *testing.T) {
	data := "avavde"
	window := map[byte]int{}
	for i := 0; i < len(data); i++ {
		value := data[i]
		window[value]++
	}
	rs, err := json.Marshal(window)
	if err == nil {
		t.Log(string(rs))
	}
}

func TestMinWindow(t *testing.T) {
	rs := minWindow("ADOBECODEBANC", "ABC")
	t.Log(rs)
}
