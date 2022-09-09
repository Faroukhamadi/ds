package hm

import (
	"math/big"
	"reflect"
)

type HashMap[K string | int, T string | int] struct {
	Entries    int
	loadFactor float64
	capacity   int
	arr        []int
}

func New[K, V string | int]() *HashMap[K, V] {
	return &HashMap[K, V]{Entries: 0, loadFactor: 0.75, capacity: 16, arr: make([]int, 16)}
}

func (h *HashMap[K, V]) Set(key K, val V) {
	var index int
	if reflect.TypeOf(key).Kind() == reflect.Int {
		index = HashInt(key, h.capacity)
	}
}

// Increase size when Entries = capacity * loadFactor
func HashInt(key, capacity int) int {
	return key % capacity
}

func Hash(key string, capacity int) int {
	hexValueBig := new(big.Int)
	hexValueBig.SetString("F0000000L", 16)

	hexValueInt := hexValueBig.Int64()

	h := 0
	for _, p := range key {
		h = (h << 4) + int(p)
		g := h & int(hexValueInt)
		if g != 0 {
			h ^= (g >> 24)
		}
		h &= ^g
	}
	return h % capacity
}
