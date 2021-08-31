package bmark

import (
	"testing"
	"time"
)

var f int

func TestFib(t *testing.T) {

	fibNumbers := []int{
		0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657,
	}

	for i := 0; i < len(fibNumbers); i++ {
		if Fib(i) != fibNumbers[i] {
			t.Error("ERROR Fib()")
			return
		}
	}
}

func BenchmarkFib(b *testing.B) {
	// Run benchmark sequancially
	for i := 0; i < b.N; i++ {
		f = Fib(10)
	}
}

func BenchmarkFibParallel(b *testing.B) {
	//Run benchmark paralally
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			f = Fib(10)
		}
	})
}

func BenchmarkFibParallelControlTimer(b *testing.B) {

	//Run benchmark paralally
	b.RunParallel(func(pb *testing.PB) {

		b.StopTimer()

		// Setting things up
		time.Sleep(time.Second * 5)

		b.StartTimer()

		for pb.Next() {
			f = Fib(10)
		}
	})
}

// Run benchmarks with multiple inputs
func benchmarkFib(val int, b *testing.B) {
	// Run benchmark sequancially
	for i := 0; i < b.N; i++ {
		f = Fib(val)
	}
}

func BenchmarkFib1(b *testing.B)  { benchmarkFib(1, b) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(2, b) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(3, b) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(10, b) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(20, b) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(40, b) }
