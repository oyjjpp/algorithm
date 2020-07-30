// 条件变量
package learn

import (
	"log"
	"sync"
)

/*
条件变量
在共享资源的状态产生变化的时候，起到通知的作用（针对生产者和消费者）

优势：提升效率

提供的三个方法
Wait 等待通知，会自动释放锁，并挂起调用者的协程；之后恢复执行，会在返回时对L加锁
除非被Signal或者Broadcast唤醒，否则Wait不会返回

Signal 单发通知 只会唤醒一个等待的协程 可以加锁，也可以不加锁
Broadcast 广播通知 会唤醒所有等待的协程，调用的时候可以加锁也可以不加锁

等待通知 需要在它基于的那个互斥锁保护下进行
发放通知 需要在对应的互斥锁解锁之后进行

互斥锁：对一个共享区域进行加锁，所有线程都是一种竞争的状态去访问


注意事项
1、条件判断使用循环
2、调用Wait()一定要加锁 防止出现死锁现象
*/

func with_cond() {
	// 共享资源
	var mailbox uint8
	var lock sync.RWMutex

	// 条件变量
	sendCond := sync.NewCond(&lock)
	recvCond := sync.NewCond(lock.RLocker())

	// sign用于传递演示完成的信号
	// sign := make(chan struct{}, 2)
	var wg sync.WaitGroup
	wg.Add(2)

	max := 5
	// 发送情报
	go func(max int) {
		defer func() {
			// sign <- struct{}{}
			wg.Done()
		}()

		for i := 1; i <= max; i++ {
			// time.Sleep(time.Millisecond * 500)
			lock.Lock()
			for mailbox == 1 {
				sendCond.Wait()
			}
			log.Printf("send [%d]: the mail box is empty.", i)
			mailbox = 1
			log.Printf("send [%d]: the letter has been sent", i)
			lock.Unlock()
			recvCond.Signal()
		}
	}(max)

	// 接收情报
	go func(max int) {
		defer func() {
			// sign <- struct{}{}
			wg.Done()
		}()

		for i := 1; i <= max; i++ {
			// time.Sleep(time.Millisecond * 500)
			lock.RLock()
			for mailbox == 0 {
				recvCond.Wait()
			}
			log.Printf("recv [%d]: the mail box is empty.", i)
			mailbox = 0
			log.Printf("recv [%d]: the letter has been received.", i)
			lock.RUnlock()
			sendCond.Signal()

		}
	}(max)

	// sycn
	wg.Wait()
	// 管道
	// <-sign
	// <-sign

	// 时间等待
	// time.Sleep(2 * time.Second)
}

func with_cond_v2() {
	shareRec := false
	wg := sync.WaitGroup{}
	wg.Add(2)
	mu := sync.Mutex{}

	lock := sync.NewCond(&mu)

	go func() {
		lock.L.Lock()

		for shareRec == false {
			log.Println("one wait")
			lock.Wait()
		}

		log.Println("one", shareRec)
		lock.L.Unlock()
		wg.Done()
	}()

	go func() {
		lock.L.Lock()

		for shareRec == false {
			log.Println("two wait")
			lock.Wait()
		}

		log.Println("two", shareRec)
		lock.L.Unlock()
		wg.Done()
	}()

	// time.Sleep(2 * time.Second)
	lock.L.Lock()
	log.Println("main goroutine ready")
	shareRec = true
	lock.Broadcast()
	log.Println("main goroutine broadcast")
	lock.L.Unlock()
	wg.Wait()
}
