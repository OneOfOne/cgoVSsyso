package main

import (
	"fmt"
	"testing"
)

func mul(a0, a1, a2, a3, a4 uint64, r *uint64)
func Mul(a0, a1, a2, a3, a4 uint64) (r uint64) {
	mul(a0, a1, a2, a3, a4, &r)
	return
}

func BenchmarkSyso(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Mul(10, 20, 30, 40, 50)
	}
}

func main() {
	b := testing.Benchmark(BenchmarkSyso)
	fmt.Println("syso:", Mul(10, 20, 30, 40, 50), b.String(), b.MemString())
}
