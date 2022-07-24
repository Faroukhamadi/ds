package main

import (
	"github.com/Faroukhamadi/ds/ll"
)

func main() {
	head := ll.New("hello")
	head = ll.Append(head, "this")
	head = ll.Append(head, "is")
	head = ll.Append(head, "my")
	head = ll.Append(head, "linked")
	head = ll.Append(head, "list")
	head = ll.Append(head, "sentence")

	head.Print()
}
