package main

import (
	"fmt"

	"github.com/Faroukhamadi/ds/ll"
)

func main() {
	head := ll.New("c")

	head = ll.Append(head, "b")
	head = ll.Append(head, "a")

	fmt.Println("BEFORE SORTING")
	head.Print()

	head = ll.Sort(head)

	fmt.Println("AFTER SORTING")
	head.Print()
}
