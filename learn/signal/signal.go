// @link https://books.studygolang.com/The-Golang-Standard-Library-by-Example/chapter16/16.03.html
package signal

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

var firstSigUsr1 = true

// signalNotify
// 信号通知
func signalNotify() {
	// 忽略 Ctrl+c
	// os.Interrupt 和 syscall.SIGINT同样含义
	signal.Ignore(os.Interrupt)

	signs := make(chan os.Signal, 1)

	// 注册信息号
	signal.Notify(signs, syscall.SIGUSR1, syscall.SIGHUP)

	go func() {
		for {
			switch <-signs {
			case syscall.SIGHUP:
				log.Println("sighub, reset sighub")
				signal.Reset(syscall.SIGHUP)
			case syscall.SIGUSR1:
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
