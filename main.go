package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	urls := []string{"https://www.zhihu.com/robots.txt", "https://www.douban.com/robots.txt"}
	resp := make(map[string]string)
	fetch_urls(urls, resp, time.Second*2)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	fmt.Println("Response:", resp)
}

// 要求
// 1、并行发送请求
// 2、要求控制全部URL获取超时时间timeout
func fetch_urls(urls []string, resp map[string]string, timeout time.Duration) {
	// 暂存结果
	var wg sync.WaitGroup
	for i := 0; i < len(urls); i++ {
		wg.Add(1)
		go func(url string, wg *sync.WaitGroup, timeout time.Duration) {
			resp[url] = curl("GET", url, timeout)
			wg.Done()
		}(urls[i], &wg, timeout)
	}
	wg.Wait()
}

func curl(method, url string, timeout time.Duration) string {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return ""
	}

	// 初始化Client 并设置超时时间
	t := timeout * time.Second
	Client := http.Client{Timeout: t}
	resp, err := Client.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		res, _ := ioutil.ReadAll(resp.Body)
		return string(res)
	}
	return ""
}
