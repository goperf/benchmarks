package bmark

import "testing"

var rs int

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rs = Add(10, 20)
	}
}
