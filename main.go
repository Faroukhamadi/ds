package main

import (
	"fmt"

	"github.com/Faroukhamadi/ds/hm"
)

func main() {
	hashMap := hm.New[int, int]()
	fmt.Println(hashMap)

	hashMap.Set(47, 2)

	fmt.Println(hashMap)

}
