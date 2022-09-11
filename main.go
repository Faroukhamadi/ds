package main

import (
	"github.com/Faroukhamadi/ds/hm"
)

func main() {
	hashMap := hm.New[int, int]()

	hashMap.Set(16, 3)
	hashMap.Set(13, 6)

	hashMap.Get(16)
	hashMap.Get(12)
	hashMap.Get(13)

}
