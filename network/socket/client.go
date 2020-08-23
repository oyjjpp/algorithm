package socket

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func Clientmain() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage:%s name", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]

	// 连接到服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8091")
	checkError(err)
	for {
		// 循环向连接中写入名字
		_, err := conn.Write([]byte(name))
		checkError(err)

		// 循环从连接中读取响应内容，没有响应时会阻塞
		response := make([]byte, 256)
		readLength, err := conn.Read(response)
		checkError(err)
		if readLength > 0 {
			log.Println("[client] server response:", string(response))
			time.Sleep(1 * time.Second)
		}
	}

}

func checkError(err error) {
	if err != nil {
		log.Fatal("fatal error:" + err.Error())
	}
}
