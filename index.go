package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/streadway/amqp"
)

const CLIENT_RSA_PUB = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDGUxT+FcvoXTYV0zweU52iL/AB
WPb3KSUu7Q1DQo6/ZbkQFsjRT/PlQHVUPvioO0O9sPAa35JxZzgE0OzWYJJIMu+g
CNAw/eXyxF79JrmDTK979lg4rzTz4NIXoo/dfsGHuG4Xr0AGeVY3aNC3KYV+HnhO
lXjYyV805CmQmxDFqQIDAQAB
-----END PUBLIC KEY-----`

// RsaVerifySHA256withRSAPSS
// 签名验证
func RsaVerifySHA256withRSAPSS(data, sign, pubkey string) error {
	clientSign, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return err
	}

	// 处理公钥
	block, _ := pem.Decode([]byte(pubkey))
	if block == nil {
		return errors.New("pubkey key error!")
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	pub := pubInterface.(*rsa.PublicKey)

	// 获取可以利用的加密hash算法
	hashed := sha256.Sum256([]byte(data))

	// 取模校验
	nBits := pub.N.BitLen()
	fmt.Println("ErrVerification", nBits, len(clientSign), (nBits+7)/8)

	//opts := rsa.PSSOptions{SaltLength: 0}
	return rsa.VerifyPSS(pub, crypto.SHA256, hashed[:], clientSign, nil)
}

func connection() *amqp.Connection {
	// 连接
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	return conn
}

func sample() {
	var wg sync.WaitGroup
	wg.Add(1)
	go sampleConsumption(&wg)

	conn := connection()
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer channel.Close()

	// 创建交换器
	if err := channel.ExchangeDeclare(
		"e1",
		"direct",
		true,
		false,
		false,
		true,
		nil); err != nil {
		panic(err)
	}

	// 创建路由器
	if _, err := channel.QueueDeclare(
		"q1",
		true,
		false,
		false,
		true,
		nil); err != nil {
		panic(err)
	}

	// 绑定队列
	if err := channel.QueueBind("q1", "q1Key", "e1", true, nil); err != nil {
		panic(err)
	}

	// mandatory true 未找到队列返回给消费者
	returnChan := make(chan amqp.Return, 0)
	channel.NotifyReturn(returnChan)

	// pushlish
	if err := channel.Publish(
		"e1",
		"q1Key",
		true,
		false,
		amqp.Publishing{
			Timestamp:   time.Now(),
			ContentType: "text/plain",
			Body:        []byte("Hello Golang and AMQP(Rabbitmq)!"),
		}); err != nil {
		panic(err)
	}

	for v := range returnChan {
		fmt.Printf("Return %#v\n", v)
	}

	wg.Wait()
}

func sampleConsumption(wg *sync.WaitGroup) {
	// 创建链接
	conn := connection()
	defer conn.Close()

	// 创建通道
	channel, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer channel.Close()

	// 消费信息
	deliveries, err := channel.Consume("q1", "any", false, false, false, true, nil)
	if err != nil {
		panic(err)
	}
	// 取一条消息
	if v, ok := <-deliveries; ok {
		if err := v.Ack(true); err != nil {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println("channel close")
	}
	wg.Done()
}

func PrintlnByteLen() {
	data := []byte("this")
	log.Println(len(data))
}
