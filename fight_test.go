package cgoVSsyso

import (
	"testing"

	cgo "github.com/OneOfOne/cgoVSsyso/cgo"
	syso "github.com/OneOfOne/cgoVSsyso/syso"
)

func BenchmarkCgo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cgo.Mul(10, 10)
	}
}

func BenchmarkSyso(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = syso.Mul(10, 10)
	}
}
