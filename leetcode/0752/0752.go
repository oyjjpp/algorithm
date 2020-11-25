package leetcode

import (
    "strings"
)

// openLock
// 打开转盘锁
func openLock(deadends []string, target string) int {
    // 初始化一个队列
    queue := []string{"0000"}
    count := []int{0}
    deads := strings.Join(deadends, ",")
    
    // 记录已经穷举过的密码，防止走回头路
    visited := map[string]bool{"0000":true}
    
    // 将当前队列中的所有节点向周围扩散
    for i:=0; i<len(queue); i++ {
        cur := queue[i]
        number := count[i]
        // 校验是否达到终点
        if strings.Contains(deads, cur){
            continue
        }
        if cur==target {
            return number
        }
        
        // 将一个节点的相邻节点加入队列
        for j:=0; j<4; j++{
            up := plusOne(cur, j)
            if _, ok := visited[up];!ok{
                queue = append(queue, up)
                count = append(count, number+1)
                visited[up] = true
            }
            
            down := minusOne(cur, j)
            if _, ok := visited[down]; !ok{
                queue = append(queue, down)
                count = append(count, number+1)
                visited[down] = true
            }
        }
    }
    return -1
}

// openLockBFS
// 问题
// 1、会走回头路
// 2、没有对deadends的处理
func openLockBFS(target string) int {
    // 初始化一个队列
    queue := []string{"0000"}
    count := []int{1}
    
    // 将当前队列中的所有节点向周围扩散
    for i:=0; i<len(queue); i++ {
        cur := queue[i]
        number := count[i]
        // 校验是否达到终点
        if cur==target {
            return number
        }
        
        // 将一个节点的相邻节点加入队列
        for j:=0; j<4; j++{
            up := plusOne(cur, j)
            queue = append(queue, up)
            count = append(count, number+1)
            
            down := minusOne(cur, j)
            queue = append(queue, down)
            count = append(count, number+1)
        }
    }
    return 0
}

// plusOne
// 将s[j]向上拨动一次
func plusOne(s string, index int) string {
    data := []byte(s)
    if data[index] == 57 {
        data[index] = 48
    } else {
        data[index]++
    }
    return string(data)
}

// minusOne
// 将 s[i] 向下拨动一次
func minusOne(s string, index int) string {
    data := []byte(s)
    if data[index] == 48 {
        data[index] = 57
    } else {
        data[index]--
    }
    return string(data)
}
