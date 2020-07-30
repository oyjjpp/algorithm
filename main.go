package main

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// 获取当前进程ID
// 获取当前线程ID
// 获取当前协程ID

func main() {
	log.Println(Goid())
	go func() {
		log.Println(Goid())
	}()
	time.Sleep(2 * time.Second)
}

func with_context() {
	ctx, cancel := context.WithCancel(context.Background())

	go work(ctx, "one")
	go work(ctx, "two")
	go work(ctx, "three")
	time.Sleep(2 * time.Second)
	log.Println("stop the goroutine")
	cancel()
	time.Sleep(time.Second * 2)
}

func with_channel() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				log.Println("one stop channel")
				return
			default:
				log.Println("one still workinf")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	log.Println("stop the goroutine")
	stop <- true
	stop <- true
	// time.Sleep(2 * time.Second)
}

func work(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			log.Println(name, "one stop channel")
			return
		default:
			log.Println(name, "one still workinf")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func Goid() int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("panic recover:panic info:%v", err)
		}
	}()

	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
