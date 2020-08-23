package common

import (
	"log"
	"math"
	"sync"
)

func isPalindrome(x int) bool {
	// 负数肯定不是回文，因为负号
	if x < 0 {
		return false
	}
	// 记录下原始数字
	originX := x

	// 反转数字
	rev := 0
	for x != 0 {
		// 每次取余数 用于反转相加
		pop := x % 10
		// 递归除以10 进行向右偏移
		x = x / 10
		rev = rev*10 + pop
	}

	// 原始数字和反转数字对比
	if originX == rev {
		return true
	}
	return false
}

// 整数反转
func reverseNumber(x int) int {
	y := 0
	for x != 0 {
		y = y*10 + x%10
		if !(-(1<<31) <= y && y <= (1<<31)-1) {
			return 0
		}
		x /= 10
	}
	return y
}

// 整数反转
func reverseNumberV2(x int) int {
	ret := 0
	for x != 0 {
		pop := x % 10
		x = x / 10
		ret = ret*10 + pop
		if ret < math.MinInt32 || ret > math.MaxInt32 {
			return 0
		}
	}
	return ret
}

// 启动两个线程, 一个输出 1,3,5,7…99, 另一个输出 2,4,6,8…100 最后 STDOUT 中按序输出 1,2,3,4,5…100？

// withCond
// 使用条件变量进行通知
func withCond() {
	var wg sync.WaitGroup
	var lock sync.Mutex

	// 定义条件变量
	var cond1 = sync.NewCond(&lock)
	var cond2 = sync.NewCond(&lock)

	wg.Add(2)

	go func() {
		defer func() {
			wg.Done()
		}()

		for i := 1; i < 100; i += 2 {
			// 获取锁
			lock.Lock()
			// 第一个直接输出
			if i != 1 {
				// 等待通知 挂起协程
				cond1.Wait()
			}
			log.Println("goroutine1->", i)
			lock.Unlock()
			// 发送通知
			cond2.Broadcast()
		}
	}()

	go func() {
		defer func() {
			wg.Done()
		}()

		for i := 2; i <= 100; i += 2 {
			lock.Lock()
			cond2.Wait()
			log.Println("goroutine2->", i)
			lock.Unlock()
			cond1.Broadcast()
		}
	}()
	wg.Wait()
}

// 启动两个线程, 一个输出 1,3,5,7…99, 另一个输出 2,4,6,8…100 最后 STDOUT 中按序输出 1,2,3,4,5…100？
// withChannel
// 使用管道方式进行通知
func withChannel() {
	var wg sync.WaitGroup
	var signal = make(chan struct{}, 1)
	var signa2 = make(chan struct{}, 1)

	wg.Add(2)
	// 输出基数
	go func() {
		defer func() {
			wg.Done()
		}()
		for i := 1; i <= 100; i += 2 {
			if i != 1 {
				// 读取数据，没有数据时阻塞
				<-signal
			}
			// <-signal
			log.Println("goroutine1->", i)
			signa2 <- struct{}{}
		}
	}()
	// 输出偶数
	go func() {
		defer func() {
			wg.Done()
		}()
		for i := 2; i <= 100; i += 2 {
			<-signa2
			log.Println("goroutine2->", i)
			signal <- struct{}{}
		}
	}()
	wg.Wait()
}
