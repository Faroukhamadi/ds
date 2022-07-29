package main

import (
	"fmt"

	"github.com/Faroukhamadi/ds/ll"
)

func main() {
	head := ll.New("abc")
	head = ll.Append(head, "bcd")

	head, err := ll.InsertPos(head, 1, "cde")
	if err != nil {
		fmt.Println(err)
	}

	head.Print()
}
