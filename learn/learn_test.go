package learn

import (
	"context"
	"errors"
	"fmt"
	"os"
	"runtime"
	"sync"
	"testing"
	"time"
	"unsafe"
)

// 使用通过复用，降低复杂对象的创建和GC代价
// 协程安全、会有锁的开销
// 生命周期收到GC影响，不适合做连接池等，需要自己管理生命周期的资源池化
func TestSyncPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			t.Log("Create a new Object.")
			return 100
		},
	}
	v := pool.Get().(int)
	t.Log(v)
	pool.Put(3)
	//pool.Put(5)
	//runtime.GC()
	v1, _ := pool.Get().(int)
	t.Log(v1)
	v2, _ := pool.Get().(int)
	t.Log(v2)
}

// 对象池概念
type ReusableObj struct{}

type ObjPool struct {
	// 用于缓冲可重用对象
	bufChan chan *ReusableObj
}

func NewObjPool(numOfObj int) *ObjPool {
	objPool := ObjPool{}
	objPool.bufChan = make(chan *ReusableObj, numOfObj)
	for i := 0; i < numOfObj; i++ {
		objPool.bufChan <- &ReusableObj{}
	}
	return &objPool
}

func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(timeout):
		return nil, errors.New("time out")
	}
}

func (p *ObjPool) ReleaseObj(obj *ReusableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}

func TestObjPool(t *testing.T) {
	pool := NewObjPool(10)
	// 溢出
	// if err := pool.ReleaseObj(&ReusableObj{}); err != nil {
	// 	t.Error(err)
	// }

	for i := 0; i < 11; i++ {
		if v, err := pool.GetObj(time.Second); err != nil {
			t.Error(err)
		} else {
			t.Logf("%T\n", v)
			// 无对象
			// if err := pool.ReleaseObj(v); err != nil {
			// 	t.Error(err)
			// }
		}
	}
}

//
func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("The Result is from %d", id)
}

// 只要有一个任务返回就算完成
func FirstResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	rs := ""
	for i := 0; i < numOfRunner; i++ {
		rs += <-ch + "\n"
	}
	return rs
}

// 等待所有结果返回
func AllResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	fmt.Println(len(ch))
	return <-ch
}

func TestFirstResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(FirstResponse())
	time.Sleep(time.Second)
	t.Log("After:", runtime.NumGoroutine())
}

type Singleton struct{}

var singInstance *Singleton
var once sync.Once

func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("Create Object")
		singInstance = new(Singleton)
	})
	return singInstance
}

func TestGetSingletonObj(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			//fmt.Println(obj)
			fmt.Printf("%x\n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}

func isChanceled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isChanceled(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cancelled")
		}(i, ctx)
	}
	cancel()
}

func isChanceledV2(ch chan struct{}) bool {
	select {
	case <-ch:
		return true
	default:
		return false
	}
}

func chacel(ch chan struct{}) {
	ch <- struct{}{}
}
func chacelV2(ch chan struct{}) {
	close(ch)
}

func TestCancelV2(t *testing.T) {
	ch := make(chan struct{})
	for i := 0; i < 5; i++ {
		go func(i int, ch chan struct{}) {
			for {
				if isChanceledV2(ch) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cancelled")
		}(i, ch)
	}
	chacelV2(ch)
	time.Sleep(time.Second)
}

func dataProduct(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		//ch <- 11
		wg.Done()
	}()

}
func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
		wg.Done()
	}()
}

// close 会发送通知
func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProduct(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Wait()
}
func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherService() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

func AsyncService() chan string {
	// 有缓冲
	//retCh := make(chan string, 1)
	// 无缓冲
	retCh := make(chan string)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited")
	}()
	return retCh
}

func AsyncServiceV2() chan string {
	retCh := make(chan string, 1)
	retCh <- "This is a message"
	return retCh
}

// 多路选择和超时设置
func TestSelect(t *testing.T) {
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case retV2 := <-AsyncServiceV2():
		t.Log(retV2)
	case <-time.After(time.Millisecond * 50):
		t.Error("time out")
	}
}

func TestChannelTask(t *testing.T) {
	ret := AsyncService()
	otherService()
	fmt.Println(<-ret)
	//time.Sleep(time.Millisecond * 300)
}

func TestCloseChannelV2(t *testing.T) {
	var wg sync.WaitGroup
	ready := make(chan int)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i, "ready.")
			<-ready
			fmt.Println(i, "running...")
		}(i)
	}
	// 保证协程都准备好
	time.Sleep(time.Second)
	// 关闭通道通知，广播
	fmt.Println("Ready?Go!")
	close(ready)

	// 向已经关闭的通道写入数据
	// output panic: send on closed channel
	// ready <- 1

	// 重复关闭通道会引发panic
	// output  panic: close of closed channel
	// close(ready)

	// 关闭空的通道会引发panic
	// output panic: close of nil channel [recovered]
	// var ch chan int
	// close(ch)
	wg.Wait()
}

func TestAsncChannel(t *testing.T) {
	ch := make(chan int, 1)
	fmt.Println("channel runing before")

	// go func() {
	// 	n := <-ch
	// 	fmt.Println(n)
	// }()
	ch <- 1
	fmt.Println(<-ch)
	//close(ch)
	//fmt.Println("channel runing")
	fmt.Println("channel runing after")
}

// channel
// 无缓冲channel 写无读会阻塞再写，读无写会阻塞在读
// 同步操作必须要有配对的goroutine出现，否则会一直阻塞

func TestCountterThreadSafe(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Log(counter)
}

func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
		// go func() {
		// 	fmt.Println(i)
		// }()
	}
	// for {
	// 	fmt.Println(runtime.NumGoroutine())
	// }
	time.Sleep(time.Millisecond * 50)
}

func TestError(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			os.Exit(-1)
		}
	}()
	fmt.Println("golang error")
	panic(errors.New("自动异常"))
}

var lessThanTwoError = errors.New("less than two error")

func DoSomeThing(p interface{}) {
	switch v := p.(type) {
	case int:
		fmt.Println("integer", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("Unknow Type")
	}
	if i, ok := p.(int); ok {
		fmt.Println("Interger", i)
		return
	}
	if i, ok := p.(string); ok {
		fmt.Println("string", i)
		return
	}
	fmt.Println("Unknow Type")

}
func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomeThing(10)
	DoSomeThing("10")
}

// 类型指针 调用方法
type Programmer interface {
	WriterHelloWorld() string
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriterHelloWorld() string {
	return `fmt.Println("hello world")`
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriterHelloWorld())
}

func Sum(ops ...int) int {
	res := 0
	for _, op := range ops {
		res += op
	}
	return res
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
}

// unicode 字符集 字符编码
// utf-8 存储实现（转换为字节序列的规则）
func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c %[1]x", c)
	}
}
func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 1
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d id not existing", n)
	}
}

func TestMapWithFuncValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1)
	t.Logf("len m1 = %d", len(m1))

	m2 := map[int]int{}
	t.Logf("len m2 = %d", len(m2))
}

func TestSlice(t *testing.T) {

}

func TestArray(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5, 6}
	slice := arr[:3]
	t.Log(len(slice), cap(slice))
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

const (
	Monday = iota
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstant(t *testing.T) {
	t.Log(Monday, Tuesday)
	a := 7
	t.Log(a&Readable, a&Writable, a&Executable)
}

func TestFirst(t *testing.T) {
	var a int = 1
	var b int = 1
	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(b)
		temp := a
		a = b
		b = temp + a
	}
}

// 类型推断
// 多个变量进行赋值
