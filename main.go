package main

import (
	"bytes"
	"flag"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strings"
	"sync"
	"syscall"
)

func main() {
	waitBug()
}

// 对共享内存保护的失误
func waitBug() {
	data := [10]int{}
	var group sync.WaitGroup
	group.Add(len(data))

	for _, p := range data {
		log.Println("当前协程数量", runtime.NumGoroutine())
		log.Println(p)
		go func(p int) {
			log.Println(p)
			defer group.Done()
		}(p)
		// group.Wait()
	}
	group.Wait()
}

// CmdParam
// 命令行参数
func CmdParam() {
	port := flag.Int("port", 80, "input port")
	flag.Parse()
	log.Println(*port)
}

// createProcess
// 创建进程
func createProcess() {
	cmd := exec.Command("tr", "a-z", "A-Z")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	log.Printf("in all caps:%q\n", out.String())
}

var firstSigUsr1 = true

// signalNotify
// 信号通知
func signalNotify() {
	// 忽略 Ctrl+c
	// os.Interrupt 和 syscall.SIGINT同样含义
	signal.Ignore(os.Interrupt)

	signs := make(chan os.Signal, 1)

	// 注册信息号
	signal.Notify(signs, syscall.SIGILL, syscall.SIGHUP)

	go func() {
		for {
			switch <-signs {
			case syscall.SIGHUP:
				log.Println("sighub, reset sighub")
				signal.Reset(syscall.SIGHUP)
			case syscall.SIGILL:
				log.Println("fisrt usr1, notify intereupt whitch had ignore!")
				ch := make(chan os.Signal, 1)
				signal.Notify(ch, os.Interrupt)
				go handlerInterrupt(ch)
			}

		}
	}()

	select {}
}

func handlerInterrupt(ch <-chan os.Signal) {
	for {
		switch <-ch {
		case os.Interrupt:
			log.Println("signal interrupt")
		}
	}
}
