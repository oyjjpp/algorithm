package leetcode

import(
    "testing"
    "log"
)

func TestAbs(t *testing.T){
    rs := abs(-10)
    log.Println(rs)

    rs = abs(20)
    log.Println(rs)
}
