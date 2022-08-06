package main

import (
	"fmt"

	"github.com/Faroukhamadi/ds/bt"
)

func main() {
	p := bt.New(1)
	q := bt.New(1)

	fmt.Println(bt.IsSameTree(p, q))
}
