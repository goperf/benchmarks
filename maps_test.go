package bmark

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

var NumItems int = 1000000

func BenchmarkMapStringKeys(b *testing.B) {
	m := make(map[string]string)
	k := make([]string, 0)

	for i := 0; i < NumItems; i++ {
		key := strconv.Itoa(rand.Intn(NumItems))
		m[key] = "value" + strconv.Itoa(i)
		k = append(k, key)
	}

	i := 0
	l := len(m)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		if _, ok := m[k[i]]; ok {
		}

		i++
		if i >= l {
			i = 0
		}
	}
}

func BenchmarkMapIntKeys(b *testing.B) {
	m := make(map[int]string)
	k := make([]int, 0)

	for i := 0; i < NumItems; i++ {
		key := rand.Intn(NumItems)
		m[key] = "value" + strconv.Itoa(i)
		k = append(k, key)
	}

	i := 0
	l := len(m)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		if _, ok := m[k[i]]; ok {
		}

		i++
		if i >= l {
			i = 0
		}
	}
}

func BenchmarkMapByteKeys(b *testing.B) {
	m := make(map[string]string)
	k := make([][]byte, 0)

	for i := 0; i < NumItems; i++ {
		key := fmt.Sprintf("%d", rand.Intn(NumItems))
		m[key] = "value" + strconv.Itoa(i)
		k = append(k, []byte(key))
	}

	i := 0
	l := len(m)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		if _, ok := m[string(k[i])]; ok {
		}

		i++
		if i >= l {
			i = 0
		}
	}
}

func BenchmarkMapInterfaceValue(b *testing.B) {
	m := make(map[int]interface{})
	k := make([]int, 0)

	for i := 0; i < NumItems; i++ {
		key := rand.Intn(NumItems)
		m[key] = struct{}{}
		k = append(k, key)
	}

	i := 0
	l := len(m)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		if _, ok := m[k[i]]; ok {
		}

		i++
		if i >= l {
			i = 0
		}
	}
}

func BenchmarkMapEmptyStructValue(b *testing.B) {
	m := make(map[int]struct{})
	k := make([]int, 0)

	for i := 0; i < NumItems; i++ {
		key := rand.Intn(NumItems)
		m[key] = struct{}{}
		k = append(k, key)
	}

	i := 0
	l := len(m)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		if _, ok := m[k[i]]; ok {
		}

		i++
		if i >= l {
			i = 0
		}
	}
}
