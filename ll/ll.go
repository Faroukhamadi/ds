package ll

import (
	"fmt"
)

type ListNode struct {
	Val  any
	Next *ListNode
}

func New(val any) *ListNode {
	return &ListNode{Val: val}
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

// func Append(head *ListNode, val any) *ListNode {
// 	newTail := New(val)
// 	cur := head
// 	for cur.Next != nil {
// 		cur = cur.Next
// 	}
// 	cur.Next = newTail
// 	return cur
// }

func Append(head *ListNode, val any) *ListNode {
	newTail := New(val)
	cur := head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = newTail
	return head
}

func Prepend(head *ListNode, val any) *ListNode {
	newHead := New(val)
	newHead.Next = head
	head = newHead
	return head
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
