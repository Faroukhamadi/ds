package main

import "github.com/Faroukhamadi/ds/ll"

func main() {
	// -----------------START----------------------------
	head := ll.New(1)
	head = ll.Append(head, 2)
	head = ll.Prepend(head, 0)
	head.Print()
	// -----------------END----------------------------
	// head.Print()

	// head := &ListNode{Val: 1}
	// second := &ListNode{Val: 2}
	// third := &ListNode{Val: 3}

	// head.Next = second
	// second.Next = third
	// third.Next = nil

	// fmt.Println(head.Val)
	// fmt.Println(second.Val)
	// fmt.Println(third.Val)
}
