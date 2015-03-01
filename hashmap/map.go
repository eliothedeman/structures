package hashmap

import (
	"hash"
	"hash/fnv"

	"github.com/eliothedeman/structures/list"
)

// Map an implimentation of a hashmap
type Map struct {
	index    []*list.LinkedList
	size     int
	capasity int
	hashFunc hash.Hash32
}

type ListElement struct {
	key string
	val interface{}
}

func NewMap() *Map {
	return &Map{
		index:    make([]*list.LinkedList, 32),
		capasity: 32,
		hashFunc: fnv.New32(),
	}
}

// hash perform a hash function on the key, and return the index it corrosponds to
func (m *Map) hash(key string, max int) int {
	m.hashFunc.Reset()
	m.hashFunc.Write([]byte(key))
	return int(m.hashFunc.Sum32()) % max

}

// grow double the capasity of the map
func (m *Map) grow() {
	m.capasity = m.capasity << 2
	old := m.index
	m.index = make([]*list.LinkedList, m.capasity)

	// for every entry that is not nil, copy it's contents to the new index
	var l *list.LinkedList
	for i := 0; i < len(old); i++ {
		l = old[i]

		if l != nil {
			// copy every element in the list
			l.Reset()
			node := l.Next()
			var le *ListElement
			for node != nil {
				le = node.(*ListElement)
				m.Put(le.key, le.val)

				node = l.Next()
			}
		}
	}
}

// Put write a key/value pair to the map
func (m *Map) Put(key string, value interface{}) {

	le := &ListElement{
		key: key,
		val: value,
	}

	i := m.hash(key, m.capasity)
	li := m.index[i]
	if li == nil {
		li = &list.LinkedList{}
		m.index[i] = li
	}
	li.Reset()
	node := li.Next()
	for node != nil {
		if node.(*ListElement).key == key {
			li.Delete(node)
		}

		node = li.Next()
	}
	m.size++
	li.Push(le)

	if m.size > m.capasity {
		m.grow()
	}
}

// Get retrieve a value given a key
func (m *Map) Get(key string) interface{} {
	i := m.hash(key, m.capasity)
	li := m.index[i]
	if li == nil {
		return nil
	}

	li.Reset()
	node := li.Next()

	for node != nil {
		le := node.(*ListElement)
		if le.key == key {
			return le.val
		}

		node = li.Next()
	}
	return nil
}
