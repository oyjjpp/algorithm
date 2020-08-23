package main

import (
	"bytes"
	"log"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"

	"github.com/algorithm/oyjjpp/network/socket"
)

func main() {
	socket.Servermain()
}

// 网络文件描述符
func netFileDesc() {
	ln, err := net.Listen("tcp", ":8091")
	if err != nil {
		log.Fatal(err)
	}
	tcpln := ln.(*net.TCPListener)
	nf, err := tcpln.File()
	ln.Close()

	fd := nf.Fd()
	log.Println(fd)

	cid, _, err := syscall.Accept(int(fd))
	if err != nil {
		log.Fatal(err)
	}
	syscall.Close(cid)
	syscall.Close(int(fd))
}

// fileDesc
// 文件描述符
func fileDesc() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fileName := path + "index.go"
	fileA, _ := os.OpenFile(fileName, os.O_RDONLY, 0666)
	log.Println(fileA.Fd())
	fileA.Close()
	fileB, _ := os.OpenFile(fileName, os.O_RDONLY, 0666)
	log.Println(fileB.Fd())
	fileB.Close()
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
	signal.Notify(signs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-signs
		log.Println(sig)
		done <- true
	}()

	log.Println("waiting singal")
	<-done
	log.Println("exting")
}

func forkProcess() {
	// file := netListener.File()
}
