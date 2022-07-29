package ll

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type ListNode[T constraints.Ordered] struct {
	Val  T
	Next *ListNode[T]
	len  int
}

func New[T constraints.Ordered](val T) *ListNode[T] {
	return &ListNode[T]{Val: val, len: 1}
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

func (l *ListNode[T]) Len() int {
	return l.len
}

func Append[T constraints.Ordered](head *ListNode[T], val T) *ListNode[T] {
	head.len++
	newTail := New(val)
	cur := head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = newTail
	return head
}

func Prepend[T constraints.Ordered](head *ListNode[T], val T) *ListNode[T] {
	head.len++
	newHead := New(val)
	newHead.len = head.len
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

func Shift[T constraints.Ordered](head *ListNode[T]) (*ListNode[T], error) {
	if head == nil {
		return nil, fmt.Errorf("[ERROR] head cannot be nil")
	}
	head = head.Next
	return head, nil
}

func Pop[T constraints.Ordered](head *ListNode[T]) (*ListNode[T], error) {
	if head == nil {
		return nil, fmt.Errorf("[ERROR] head cannot be nil")
	}
	cur := head
	for cur.Next.Next != nil {
		cur = cur.Next
	}
	cur.Next = nil
	return head, nil
}

func InsertPos[T constraints.Ordered](head *ListNode[T], pos int, val T) (*ListNode[T], error) {
	if pos < 1 || pos > head.len+1 {
		return nil, fmt.Errorf("[ERROR] entered position is not valid")
	}
	head.len++
	if pos == 1 {
		head = Prepend(head, val)
		return head, nil
	} else if pos == head.len {
		head = Append(head, val)
		return head, nil
	}
	cur := head
	for i := 0; i < pos; i++ {
		if i == pos-2 {
			newNode := New(val)
			next := cur.Next
			cur.Next = newNode
			newNode.Next = next
			break
		} else {
			cur = cur.Next
		}
	}
	return head, nil
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
