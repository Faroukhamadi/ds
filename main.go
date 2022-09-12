package main

import (
	"fmt"

	"github.com/Faroukhamadi/ds/hm"
)

func main() {
	hashMap := hm.New[int, int]()

	hashMap.Set(16, 3)

	fmt.Println(hashMap.Get(16))

	fmt.Println(hashMap.Get(12))

	hashMap.Set(13, 6)

	fmt.Println(hashMap.Get(13))

	hashMap.Set(16, 2)

	fmt.Println(hashMap.Get(16))

	fmt.Println(hashMap.Get(69))

	fmt.Println(hashMap.Entries)
}
