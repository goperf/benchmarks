package bmark

import (
	"sync"
	"testing"
)

type BigStruct struct {
	Field1 int
	Field2 string
	Field3 []float64
	Field4 map[string]interface{}
	Field5 *NestedStruct
}

type NestedStruct struct {
	NestedField1 bool
	NestedField2 []string
}

var pool = &sync.Pool{
	New: func() interface{} {
		return new(BigStruct)
	},
}

var st *BigStruct

func BenchmarkWithoutPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		st = new(BigStruct)
		st.Field1 = 1
		st.Field2 = "test"
	}
}

func BenchmarkWithPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		st = pool.Get().(*BigStruct)
		st.Field1 = 1
		st.Field2 = "test"

		// Returing the allocation to the pool once the task is done.
		pool.Put(st)
	}
}
