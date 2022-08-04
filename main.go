package main

import (
	"github.com/Faroukhamadi/ds/bt"
)

func main() {
	node := bt.New(1)
	node.Left = bt.New(2)
	node.Right = bt.New(3)
	node.Left.Left = bt.New(4)
	node.Right.Right = bt.New(5)

	bt.Print(node, 0)

	bt.Invert(node)

	bt.Print(node, 0)

}
