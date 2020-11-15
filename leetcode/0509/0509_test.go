package leetcode

import (
	"log"
	"testing"
)

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(20)
	}
}

func BenchmarkFibV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibV2(20)
	}
}

func BenchmarkFibV3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibV3(20)
	}
}

func BenchmarkFibV4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibV4(20)
	}
}

func TestFibV2(t *testing.T) {
	rs := fibV2(2)
	log.Println(rs)

	data := fib(2)
	log.Println(data)
}
