package hm

import (
	"log"
	"math/big"
	"reflect"

	"github.com/Faroukhamadi/ds/ll"
)

type HashMap[K string | int, V string | int] struct {
	Entries    int
	LoadFactor float64
	Capacity   int
	// IMPORTANT: problem with this is that the key and value have to be the same
	// Change node to ListNode[K, V]
	Arr []*ll.ListNode[K, V]
}

func New[K string | int, V string | int]() *HashMap[K, V] {
	return &HashMap[K, V]{Entries: 0, LoadFactor: 0.75, Capacity: 16, Arr: make([]*ll.ListNode[K, V], 16)}
}

func HashInt(key interface{}, capacity int) int {
	return key.(int) % capacity
}

func HashStr(key interface{}, capacity int) int {
	hexValueBig := new(big.Int)
	hexValueBig.SetString("F0000000L", 16)

	hexValueInt := hexValueBig.Int64()

	h := 0
	for _, p := range key.(string) {
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
	var index int
	if reflect.TypeOf(key).Kind() == reflect.String {
		index = HashStr(key, h.Capacity)
	} else if reflect.TypeOf(key).Kind() == reflect.Int {
		index = HashInt(key, h.Capacity)
	}

	log.Println("index is: ", index)

	if h.Arr[index] == nil {
		h.Arr[index] = ll.New[K](val)
		h.Arr[index].Key = key
		h.Arr[index].Print()
	}
}

// Increase size when Entries = capacity * loadFactor
