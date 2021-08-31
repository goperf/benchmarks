package bmark

import "testing"

func TestEscapeAnalysis(t *testing.T) {
	EscapeAnalysis()
}

func BenchmarkEscapeAnalysis(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EscapeAnalysis()
	}
}
