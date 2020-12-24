package leetcode

import (
    "testing"
    "reflect"
)

func TestIsMatch(t *testing.T){
    // "aa" "a" false
    // "aab" "c*a*b" true
    rs := isMatch("aab", "c*a*b")
    t.Log(rs)
}

func TestStringPtr(t *testing.T){
    str := "this is a message"
    m := &str
    
    t.Log(reflect.ValueOf(m).Kind())
    t.Log(reflect.ValueOf(*m).Kind())
    t.Log((*m)[1])
}
