package main

import "fmt"

func main() {
	param := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	for item, v := range param {
		fmt.Println(item)
		fmt.Println(v)
	}

	fmt.Println("length=", len(param))
	mapData := map[string]string{}
	fmt.Println("length=", len(mapData))
}

func TestRs() (string, int) {
	return "abc", 12
}
