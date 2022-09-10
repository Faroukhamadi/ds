package hm

import (
	"math/big"
)

type String string
type Int int

type Key interface {
	*String | *Int
	Hash(int) int
}

type HashMap[K Key, T string | int] struct {
	entries    int
	loadFactor float64
	capacity   int
	arr        []int
}

func New[K Key, V string | int]() *HashMap[K, V] {
	return &HashMap[K, V]{entries: 0, loadFactor: 0.75, capacity: 16, arr: make([]int, 16)}
}

func Hash[T Key](key T, capacity int) int {
	return key.Hash(capacity)
}

func (key *Int) Hash(capacity int) int {
	return int(*key) % capacity
}

func (key *String) Hash(capacity int) int {
	hexValueBig := new(big.Int)
	hexValueBig.SetString("F0000000L", 16)

	hexValueInt := hexValueBig.Int64()

	h := 0
	for _, p := range *key {
		h = (h << 4) + int(p)
		g := h & int(hexValueInt)
		if g != 0 {
			h ^= (g >> 24)
		}
		h &= ^g
	}
	return h % capacity
}

func (h *HashMap[K, V]) Set(key K, val V) {
	index := key.Hash(h.capacity)
}

// Increase size when Entries = capacity * loadFactor
