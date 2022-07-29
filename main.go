package main

import (
	"log"

	"github.com/Faroukhamadi/ds/ll"
)

func main() {
	head := ll.New("a")
	head = ll.Append(head, "b")
	head = ll.Append(head, "c")
	head = ll.Append(head, "d")
	head = ll.Append(head, "e")

	head.Print()
	head, err := ll.Pop(head)
	if err != nil {
		log.Println(err)
	}
	head.Print()
}
