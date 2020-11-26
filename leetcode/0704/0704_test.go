package leetcode

import (
    "testing"
)

func TestSearch(t *testing.T){
    data := []int{-1,0,3,5,9,12}
    index := search(data, 9)
    if index != 4 {
        t.Logf("期望%d, 实际%d\n", 4, index)
    }
    
    index = search(data, 2)
    if index != -1 {
        t.Logf("期望%d, 实际%d\n", -1, index)
    }
}
