package channel

import (
	"context"
	"errors"
	"fmt"
	"html"
	"log"
	"math/rand"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

func timeWork() <-chan int {
	time.Sleep(2 * time.Second)
	done := make(chan int, 1)
	done <- 1
	return done
}

// withChannelTime
// 通过channel进行超时控制
func withChannelTime() {
	for {
		select {
		case <-time.After(1 * time.Second):
			log.Println("timeout")
		case work := <-timeWork():
			log.Println("sucess", work)
		default:
			log.Println("default")
		}
	}
}

// withNoDataChanel
// 无缓冲通道阻塞
func withNoDataChanel() {
	data := make(chan int)
	log.Println("从非缓冲通道读取1")
	<-data
	log.Println("从非缓冲通道读取2")
}

// readNoDataFromNoBufChannelWithSelect
// 非缓冲通道无数据读取
func readNoDataFromNoBufChannelWithSelect() {
	data := make(chan int)
	if v, err := readWithSelect(data); err != nil {
		log.Println(err)
	} else {
		log.Println("read", v)
	}
}

func readWithSelect(ch chan int) (int, error) {
	select {
	case x := <-ch:
		return x, nil
	default:
		return 0, errors.New("channel has no data")
	}
}

func task() {
	log.Println("任务")
}

// withEmptyChannel
// 使用空的通道退出select
func withEmptyChannel() {
	work := asChan(1, 2)
	for {
		select {
		case rs, ok := <-work:
			if !ok {
				work = nil
				//break
			}
			log.Println(rs)
			time.Sleep(2 * time.Second)
		default:
			log.Println("无选择")
			time.Sleep(2 * time.Second)
		}
	}
}
func asChan(vs ...int) chan int {
	ch := make(chan int)
	go func() {
		for _, v := range vs {
			ch <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(ch)
	}()
	return ch
}

// writeWithSelect
// 通过select进行写入操作
func writeWithSelect(ch chan int) error {
	select {
	case ch <- 1:
		return nil
	default:
		return errors.New("The pipe is full")
	}
}

func readWithSelectTime(ch chan int) (int, error) {
	timeout := time.NewTimer(time.Microsecond * 500)
	select {
	case x := <-ch:
		return x, nil
	case <-timeout.C:
		return 0, errors.New("read timeout")
	}
}

func writeWithSelectTime(ch chan int) error {
	timeout := time.NewTimer(time.Microsecond * 500)
	select {
	case ch <- 1:
		return nil
	case <-timeout.C:
		return errors.New("write time out")
	}
}

// readNoDataFromBufChannelWithSelect
// 缓冲通道无数据读取
func readNoDataFromBufChannelWithSelect() {
	data := make(chan int, 1)
	if v, err := readWithSelect(data); err != nil {
		log.Println(err)
	} else {
		log.Println("read", v)
	}
}

/*
select是执行选择操作的一个结构，包含一组case语句，会执行其中无阻塞的某一个case，如果全部阻塞了，那就等待其中一个
不阻塞进而继续执行；还有一个defalut语句，default语句永远不会阻塞，可以借助他实现无阻塞操作
*/

func testChannel() {
	data := make(chan int, 3)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer func() {
			wg.Done()
		}()
		for {
			if value, ok := <-data; ok {
				log.Println("消费", value)
			} else {
				return
			}
		}
	}()
	go func() {
		defer func() {
			wg.Done()
		}()
		for i := 0; i < 3; i++ {
			data <- i
			log.Println("生产", i)
		}
		// close(data)
	}()

	wg.Wait()
}

// TrimHtml
// 去掉HTML标签
// @param src 原字符串
func TrimHtml(src string, newline bool) string {
	// 将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	// 去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")

	// 去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	// 去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "")

	// 去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	if newline {
		src = re.ReplaceAllString(src, "\n")
	} else {
		src = re.ReplaceAllString(src, "  ")
	}
	src = html.UnescapeString(src)
	return strings.TrimSpace(src)
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
