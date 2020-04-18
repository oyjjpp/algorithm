package main

import (
	"fmt"
	"time"
)

func main() {
	theMine := []string{"ore1", "ore2", "ore3"}
	oreChan := make(chan string)

	// 协程1
	go func(item []string) {
		for _, v := range item {
			oreChan <- v
			fmt.Println("send data")
		}
	}(theMine)

	// 协程2
	go func() {
		for i := 0; i < 3; i++ {
			foundOne := <-oreChan
			fmt.Println("Miner:Received " + foundOne + " from finder")
		}
	}()
	<-time.After(time.Second * 2)
}
