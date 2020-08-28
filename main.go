package main

import (
	"bytes"
	"database/sql"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"

	"github.com/pkg/errors"
)

func getSql() error {
	return errors.Wrap(sql.ErrNoRows, "GetSql failed")
}

func call() error {
	return errors.WithMessage(getSql(), "bar failed")
}

func withStack() {
	err := call()
	if errors.Cause(err) == sql.ErrNoRows {
		log.Printf("data not found, %v\n", err)
		log.Printf("%+v\n", err)
		return
	}
	if err != nil {
		log.Printf("got err,  %+v\n", err)
	}
}

func main() {
	withStack()
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
