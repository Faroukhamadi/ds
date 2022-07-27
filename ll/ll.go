package ll

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type ListNode[T constraints.Ordered] struct {
	Val  T
	Next *ListNode[T]
}

func New[T constraints.Ordered](val T) *ListNode[T] {
	return &ListNode[T]{Val: val}
}

func (l *ListNode[T]) Print() {
	fmt.Println("LINKED LIST")
	cur := l
	for cur != nil {
		fmt.Printf("%v -> ", cur.Val)
		cur = cur.Next
	}
	fmt.Println("nil")
}

func (l *ListNode[T]) IsInterfaceNil() bool {
	return l == nil
}

func Append[T constraints.Ordered](head *ListNode[T], val T) *ListNode[T] {
	newTail := New(val)
	cur := head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = newTail
	return head
}

func Prepend[T constraints.Ordered](head *ListNode[T], val T) *ListNode[T] {
	newHead := New(val)
	newHead.Next = head
	head = newHead
	return head
}

func Reverse[T constraints.Ordered](head *ListNode[T]) *ListNode[T] {
	if head == nil {
		return nil
	}
	var prev *ListNode[T]
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

func Sort[T constraints.Ordered](head *ListNode[T]) *ListNode[T] {
	if head == nil || head.Next == nil {
		return head
	}

	mid := getMid(head)
	left := Sort(head)
	right := Sort(mid)
	return merge(left, right)
}

func getMid[T constraints.Ordered](head *ListNode[T]) *ListNode[T] {
	var midPrev *ListNode[T] = nil
	var slow, fast = head, head
	for fast != nil && fast.Next != nil {
		midPrev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	var mid = midPrev.Next
	midPrev.Next = nil
	return mid
}

func merge[T constraints.Ordered](l1, l2 *ListNode[T]) *ListNode[T] {
	var head = &ListNode[T]{}
	var curr = head

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}

	if l1 != nil {
		curr.Next = l1
	}

	if l2 != nil {
		curr.Next = l2
	}

	return head.Next
}
