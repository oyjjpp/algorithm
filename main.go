package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"

	"github.com/oyjjpp/algorithm/network/socket"
)

func main() {
	// socket.Servermain()
	socket.Clientmain()
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

// signalNotify
// 信号通知
func signalNotify() {
	signs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// 注册信息号
	signal.Notify(signs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGUSR2)

	go func() {
		sig := <-signs
		log.Println(sig)
		done <- true
	}()

	log.Println("waiting singal")
	<-done
	log.Println("exting")
}
