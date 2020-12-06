package leetcode 

import (
    "testing"
    "fmt"
)

func TestGenerate(t *testing.T){
    rs := generate(5)
    t.Log(rs, len(rs))
    fmt.Printf("\n\t")
    for i:=0;i<len(rs);i++{
        for j:=0;j<len(rs[i]);j++{
            fmt.Printf("%d", rs[i][j])
        }
        fmt.Printf("\n\t")
    }
}
