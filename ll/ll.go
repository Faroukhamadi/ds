package ll

import (
	"fmt"
	"log"
)

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

func (l *ListNode) IsInterfaceNil() bool {
	if l == nil {
		return true
	}
	return false
}

func Append(head *ListNode, val any) *ListNode {
	newTail := New(val)
	cur := head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = newTail
	return cur
}

func Prepend(head *ListNode, val any) *ListNode {
	log.Println("wohooo we are inside prepend function")
	newHead := New(val)
	newHead.Next = head
	head = newHead
	return head
}

// func (l *ListNode) Prepend(val any) {
// 	head := New(val)
// 	next := l.Next
// 	l.Next = head
// 	head.Next = next
// 	swap(&l.Val, &head.Val)
// }

func swap[T any, PtrT *T](a, b PtrT) {
	fmt.Println("values before swap: ", *a, *b)
	aux := *a
	*a = *b
	*b = aux
	fmt.Println("values after swap: ", *a, *b)
}

func Reverse(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}
