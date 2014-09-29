package main

/*
#cgo CFLAGS: -march=native -O3
#include <stdint.h>
uint64_t mul(uint64_t a, uint64_t b) {
	return (a * b);
}
*/
import "C"
import (
	"fmt"
	"testing"
)

func Mul(a, b uint64) uint64 {
	return uint64(C.mul(C.uint64_t(a), C.uint64_t(b)))
}

func BenchmarkCgo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Mul(10, 10)
	}
}

func main() {
	b := testing.Benchmark(BenchmarkCgo)
	fmt.Println("syso:", b.String(), b.MemString())
}
