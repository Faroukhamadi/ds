package main

import "fmt"

type ListNode struct {
	Val  any
	Next *ListNode
}

func New(T any) *ListNode {
	return &ListNode{Val: T}
}

func (l *ListNode) Print() {
	fmt.Println("LINKED LIST")
	cur := l
	for cur != nil {
		fmt.Printf("%v -> ", cur.Val)
		cur = cur.Next
	}
	fmt.Println("nil")
}

func (l *ListNode) Append(val any) {
	newTail := New(val)
	cur := l
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = newTail
}

func swap[T any, PtrT *T](a, b PtrT) {
	fmt.Println("values before swap: ", *a, *b)
	aux := *a
	*a = *b
	*b = aux
	fmt.Println("values after swap: ", *a, *b)
}

func (l *ListNode) Prepend(val any) {
	head := New(val)
	next := l.Next
	l.Next = head
	head.Next = next
	swap(&l.Val, &head.Val)
}

func main() {
	// -----------------START----------------------------
	head := New(1)
	second := New(2)
	third := New(3)

	head.Next = second
	second.Next = third
	third.Next = nil

	head.Append(4)
	head.Prepend(0)
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
