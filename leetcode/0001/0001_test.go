package leetcode

import (
    "log"
    "testing"
)

type paramType struct{
    num []int
    target int
}

type answerType []int

type question struct{
    param   paramType
    answer answerType
}

func TestTwoSum(t *testing.T){
    params := []question{
        {
            paramType{[]int{3,2,4},6},
            answerType([]int{0,1}),
        },
    }

    for _, item := range params {
        value := item.param
        log.Printf("【input】:%v       【output】:%v\n", value, twoSum(value.num, value.target))
    }
}

