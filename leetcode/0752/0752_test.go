package leetcode

import (
    "testing"
    "log"
)

func TestString(t *testing.T){
    data := []string{"0123","1111"}
    log.Println(data[0][1])
}

func TestPlusOne(t *testing.T){
    data := []string{"0123","1111"}
    log.Println(data)
    for index, value := range data{
        data[index] = plusOne(value, index)
    }
    log.Println(data)
}

func TestMinusOne(t *testing.T){
    data := []string{"0123","1111"}
    log.Println(data)
    for index, value := range data{
        data[index] = minusOne(value, index)
    }
    log.Println(data)
}

func TestOpenLockBFS(t *testing.T){
    rs := openLockBFS("2000")
    log.Println(rs)
}

func TestOpenLock(t *testing.T){
    // ["8887","8889","8878","8898","8788","8988","7888","9888"]
    // "8888"
    // "0201","0101","0102","1212","2002"
    // "0202"
    rs := openLock([]string{"8887","8889","8878","8898","8788","8988","7888","9888"}, "8888")
    log.Println(rs)
}
