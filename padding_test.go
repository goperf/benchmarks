package bmark

import "testing"

var t interface{}

func BenchmarkNotPadded(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t = plainStruct{}
	}
}

func BenchmarkPadded(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t = optimizedStruct{}
	}
}
