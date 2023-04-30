package main

import (
	"fmt"
	"log"
)

func main() {
	data := []int{8, 3, 5, 1, 9, 12, 5}
	log.Println(data)

	s := Solution{}
	nums := []int{5, 2, 6, 1}
	res := s.countSmaller(nums)
	fmt.Println(res)
	log.Println(data)
}

type Pair struct {
	val int
	id  int
}

type Solution struct {
	temp  []Pair
	count []int
}

func (s *Solution) countSmaller(nums []int) []int {
	n := len(nums)
	s.count = make([]int, n)
	s.temp = make([]Pair, n)
	arr := make([]Pair, n)
	for i := 0; i < n; i++ {
		arr[i] = Pair{val: nums[i], id: i}
	}
	s.sort(arr, 0, n-1)
	res := make([]int, n)
	for i, c := range s.count {
		res[i] = c
	}
	return res
}

func (s *Solution) sort(arr []Pair, lo int, hi int) {
	if lo == hi {
		return
	}
	mid := lo + (hi-lo)/2
	s.sort(arr, lo, mid)
	s.sort(arr, mid+1, hi)
	s.merge(arr, lo, mid, hi)
}

func (s *Solution) merge(arr []Pair, lo int, mid int, hi int) {
	for i := lo; i <= hi; i++ {
		s.temp[i] = arr[i]
	}
	i, j := lo, mid+1
	for p := lo; p <= hi; p++ {
		if i == mid+1 {
			arr[p] = s.temp[j]
			j++
		} else if j == hi+1 {
			arr[p] = s.temp[i]
			s.count[arr[p].id] += j - mid - 1
			i++
		} else if s.temp[i].val > s.temp[j].val {
			arr[p] = s.temp[j]
			j++
		} else {
			arr[p] = s.temp[i]
			s.count[arr[p].id] += j - mid - 1
			i++
		}
	}
}
