package hashmap

import "github.com/eliothedeman/structures/list"

import (
	"sync/atomic"
)

// Map an implimentation of a hashmap
type Map struct {
	index []*list.LinkedList
	size  int64
}

// hash perform a hash function on the key, and return the index it corrosponds to
func (m *Map) hash(key string) int {
}

// Size atomically return the size of the map it it's current state
func (m *Map) Size() int64 {
	return atomic.LoadInt64(&m.size)
}
