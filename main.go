package main

import (
	"fmt"

	"github.com/Faroukhamadi/ds/hm"
	linkedlist "github.com/Faroukhamadi/ds/linked_list"
)

func main() {
	hashMap := hm.New[int, int]()

	for i := 1; i < 120; i++ {
		hashMap.Set(i, i)
	}

	l := &linkedlist.ListNode[int]{Val: 1}
	second := &linkedlist.ListNode[int]{Val: 2}
	third := &linkedlist.ListNode[int]{Val: 3}

	l.Next = second
	second.Next = third
	third.Next = l

	if l.HasCycle() {
		fmt.Println("The linked list has a cycle")
	} else {
		fmt.Println("The linked list doesn't have a cycle")
	}

}
