package ll

import (
	"fmt"
)

type ListNode[K string | int, V string | int] struct {
	Val  V
	Key  K
	Next *ListNode[K, V]
	len  int
}

func New[K string | int, V string | int](val V) *ListNode[K, V] {
	return &ListNode[K, V]{Val: val, len: 1}
}

func (l *ListNode[K, V]) Print() {
	fmt.Println("LINKED LIST")
	cur := l
	for cur != nil {
		fmt.Printf("%v -> ", cur.Val)
		cur = cur.Next
	}
	fmt.Println("nil")
}

func (l *ListNode[K, V]) IsInterfaceNil() bool {
	return l == nil
}

func (l *ListNode[K, V]) Len() int {
	return l.len
}

func Append[K string | int, V string | int](head *ListNode[K, V], val V) *ListNode[K, V] {
	head.len++
	newTail := New[K](val)
	cur := head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = newTail
	return head
}

func Prepend[K string | int, V string | int](head *ListNode[K, V], val V) *ListNode[K, V] {
	head.len++
	newHead := New[K](val)
	newHead.len = head.len
	newHead.Next = head
	head = newHead
	return head
}

func Reverse[K string | int, V string | int](head *ListNode[K, V]) *ListNode[K, V] {
	if head == nil {
		return nil
	}
	var prev *ListNode[K, V]
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

func Sort[K string | int, V string | int](head *ListNode[K, V]) *ListNode[K, V] {
	if head == nil || head.Next == nil {
		return head
	}

	mid := getMid(head)
	left := Sort(head)
	right := Sort(mid)
	return merge(left, right)
}

func Shift[K string | int, V string | int](head *ListNode[K, V]) (*ListNode[K, V], error) {
	if head == nil {
		return nil, fmt.Errorf("[ERROR] head cannot be nil")
	}
	head = head.Next
	return head, nil
}

func RemoveDuplicates[K string | int, V string | int](head *ListNode[K, V]) (*ListNode[K, V], error) {
	if head == nil {
		return nil, fmt.Errorf("[ERROR] head cannot be nil")
	}
	cur := head
	var next_next *ListNode[K, V]

	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			next_next = cur.Next.Next
			cur.Next = next_next
		} else {
			cur = cur.Next
		}
	}
	return head, nil
}

func Pop[K string | int, V string | int](head *ListNode[K, V]) (*ListNode[K, V], error) {
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

func DeletePos[K string | int, V string | int](head *ListNode[K, V], pos int) (*ListNode[K, V], error) {
	if pos < 1 || pos > head.len {
		return nil, fmt.Errorf("[ERROR] entered position is not valid")
	}
	head.len++
	if pos == 1 {
		head, _ := Shift(head)
		return head, nil
	} else if pos == head.len {
		head, _ := Pop(head)
		return head, nil
	}
	cur := head
	for i := 0; i < pos; i++ {
		if i == pos-2 {
			cur.Next = cur.Next.Next
			break
		} else {
			cur = cur.Next
		}
	}
	return head, nil
}

func InsertPos[K string | int, V string | int](head *ListNode[K, V], pos int, val V) (*ListNode[K, V], error) {
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
			newNode := New[K](val)
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

func getMid[K string | int, V string | int](head *ListNode[K, V]) *ListNode[K, V] {
	var midPrev *ListNode[K, V] = nil
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

func merge[K string | int, V string | int](l1, l2 *ListNode[K, V]) *ListNode[K, V] {
	var head = &ListNode[K, V]{}
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
