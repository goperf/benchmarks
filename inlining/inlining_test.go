package inlining

import "testing"

var rs int

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rs = Add(10, 20)
	}
}

func BenchmarkAddNoInlined(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rs = AddNoInlined(10, 20)
	}
}
