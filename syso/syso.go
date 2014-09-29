package main

import (
	"fmt"
	"testing"
)

func mul(a, b uint64, r *uint64)
func Mul(a, b uint64) (r uint64) {
	mul(a, b, &r)
	return
}

func BenchmarkSyso(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Mul(10, 10)
	}
}

func main() {
	b := testing.Benchmark(BenchmarkSyso)
	fmt.Println("syso:", b.String(), b.MemString())
}
