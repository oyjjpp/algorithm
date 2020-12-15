package leetcode

import(
    "testing"
)

func TestLongestCommonSubsequence(t *testing.T){
    text1 := "abcde" 
    text2 := "ace" 
    rs := longestCommonSubsequence(text1, text2)
    t.Log(rs)
}
