package fib

import (
	"testing"
)

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(10)
	}
}

func BenchmarkFibOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibOptimized(10)
	}
}
