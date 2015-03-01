package hashmap

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	m := NewMap()
	h1 := m.hash("hello", 1024)
	h2 := m.hash("hello", 1024)
	if h1 != h2 {
		t.Fail()
	}
}

func TestGet(t *testing.T) {
	m := NewMap()

	for i := 0; i < 100; i++ {
		k := fmt.Sprintf("%d", i)
		m.Put(k, k)
	}

	for i := 0; i < 100; i++ {
		k := fmt.Sprintf("%d", i)
		if k != m.Get(k) {
			t.Fail()
		}
	}
}

func BenchmarkHash8(b *testing.B) {
	m := NewMap()
	for i := 0; i < b.N; i++ {
		m.hash("12345678", 1024)
	}
}
func BenchmarkHash16(b *testing.B) {
	m := NewMap()
	for i := 0; i < b.N; i++ {
		m.hash("8888888888888888", 1024)
	}
}

func BenchmarkGet(b *testing.B) {
	m := NewMap()
	size := 100
	for i := 0; i < size; i++ {
		k := fmt.Sprintf("%d", i)
		m.Put(k, k)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Get(fmt.Sprintf("%d", i%size))
	}
}
