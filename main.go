package main

import (
	"fmt"

	"github.com/Faroukhamadi/ds/bt"
)

func main() {

	p := bt.New(1)
	q := bt.New(2)
	x := bt.New(3)
	y := bt.New(4)
	a := bt.New(5)
	b := bt.New(6)

	p.Left = q
	q.Left = x
	x.Left = y
	y.Left = a
	a.Right = b
	b.Left = nil

	bt.Print(p, 0)

	fmt.Println("max depth is: ", bt.MaxDepth(p))
}
