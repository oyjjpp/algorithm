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
	// rs := minWindow("ADOBECODEBANC", "ABC")
    // rs := minWindow("a", "a")
    rs := minWindow("ADOBEC", "ABC")
	t.Log(rs)
}
