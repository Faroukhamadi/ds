package main

import (
	"fmt"

	"github.com/Faroukhamadi/ds/ll"
)

func main() {
	head := ll.New("abc")

	head, err := ll.InsertPos(head, 1, "bcd")
	if err != nil {
		fmt.Println(err)
	}

	head.Print()
}
