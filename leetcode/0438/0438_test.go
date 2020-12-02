package leetcode

import (
    "testing"
)

func TestFindAnagrams(t *testing.T){
    // cbaebabacd
    data := findAnagrams("cbaebabacd", "abc")
    t.Log(data)
}
