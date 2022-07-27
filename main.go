package main

import (
	"github.com/Faroukhamadi/ds/ll"
)

func main() {
	head := ll.New("abc")
	head = ll.Append(head, "bcd")
	head = ll.Append(head, "cde")
	head = ll.Prepend(head, "prepended")

	// head.Print()
}
