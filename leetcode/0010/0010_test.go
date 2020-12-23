package leetcode

import (
    "testing"
    "reflect"
)

func TestIsMatch(t *testing.T){
    rs := false 
    rs=isMatch("aa", "a")
    t.Log(rs)
}

func TestStringPtr(t *testing.T){
    str := "this is a message"
    m := &str
    t.Log(reflect.ValueOf(m).Kind())
    t.Log(*m[1])
}
