package socket

import (
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "8091")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		// 循环接收客户端的连接，没有连接时会阻塞，出错则跳出循环
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			break
		}
		log.Println("[server] accpet new connectiont.")

		// 启动一个goroutine 处理连接
		go handler(conn)
	}
}

func handler(conn net.Conn) {
	defer conn.Close()

	for {
		// 循环从连接中读取请求内容，没有请求时会阻塞，出错则跳出循环
		request := make([]byte, 128)
		readLength, err := conn.Read(request)
		if err != nil {
			log.Println(err)
			break
		}

		if readLength == 0 {
			log.Println(err)
			break
		}

		log.Println("[server] request from", string(request))
		conn.Write([]byte("hello" + string(request) + ", time" + time.Now().Format("2006-01-02 15:04:05")))
	}
}
