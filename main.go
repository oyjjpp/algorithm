package main

import "fmt"

func main() {
	mychannel := make(chan int, 10)
	for i := 0; i < 10; i++ {
		mychannel <- i
	}
	//close(mychannel)
	fmt.Println("data length:", len(mychannel))
	for v := range mychannel {
		fmt.Println(v)
	}
	fmt.Println("data length:", len(mychannel))
}
