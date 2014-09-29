package main

/*
#cgo CFLAGS: -march=native -O3
#include <stdint.h>
uint64_t mul(uint64_t a0, uint64_t a1, uint64_t a2, uint64_t a3, uint64_t a4) {
	return a0 * a1 * a2 * a3 * a4;
}
*/
import "C"
import (
	"fmt"
	"testing"
)

func BenchmarkCgo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = C.mul(10, 20, 30, 40, 50)
	}
}

func main() {
	b := testing.Benchmark(BenchmarkCgo)
	fmt.Println("syso:", C.mul(10, 20, 30, 40, 50), b.String(), b.MemString())
}
