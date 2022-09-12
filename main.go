package main

import (
	"github.com/Faroukhamadi/ds/hm"
)

func main() {
	hashMap := hm.New[int, int]()

	for i := 1; i < 120; i++ {
		hashMap.Set(i, i)
	}
}
