package leetcode

import(
    "testing"
    "sort"
)

func TestEnvelopesData(t *testing.T){
    data := [][]int{
        {5,4},
        {6,4},
        {6,7},
        {2,3},   
    }
    t.Log(data)
    
    content := envelopesData{data}
    sort.Sort(content)
    t.Log(content.data)
}

func TestMaxEnvelopes(t *testing.T){
    data := [][]int{
        {5,4},
        {6,4},
        {6,7},
        {2,3},   
    }
    rs := maxEnvelopes(data)
    t.Log(rs)
}
