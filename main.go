package main

import "fmt"

func main() {
	item := []int{2, 1}
	rs := searchTwo(item)
	fmt.Println(rs)
}

func searchTwo(item []int) int {
	if len(item) < 2 {
		return 0
	}

	one := item[0]
	two := item[1]
	if two > one {
		one, two = two, one
	}
	for i := 2; i < len(item); i++ {
		if item[i] > one {
			two = one
			one = item[i]
		} else if item[i] > two {
			two = item[i]
		}
	}
	return two
}
