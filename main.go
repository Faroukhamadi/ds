package main

import (
	"log"

	"github.com/Faroukhamadi/ds/ll"
)

func main() {
	head := ll.New("b")

	head = ll.Append(head, "a")
	head = ll.Append(head, "a")
	head = ll.Append(head, "d")
	head = ll.Append(head, "d")
	head = ll.Append(head, "c")
	head = ll.Append(head, "c")

	head.Print()

	head = ll.Sort(head)

	head = nil
	head, err := ll.RemoveDuplicates(head)
	if err != nil {
		log.Fatal(err)
	}

	head.Print()
}
