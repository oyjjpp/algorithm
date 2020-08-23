package socket

import (
	"log"
	"net"
	"os"
	"syscall"
)

// NetFileDesc
// 网络文件描述符
func NetFileDesc() {
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

// FileDesc
// 文件描述符
func FileDesc() {
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
