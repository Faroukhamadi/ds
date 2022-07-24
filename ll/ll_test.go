package ll_test

import (
	"testing"

	"github.com/Faroukhamadi/ds/ll"
)

func TestNew(t *testing.T) {
	t.Run("Right value gets assigned to node", func(t *testing.T) {
		head := ll.New(1)

		AssertEqual(t, head.Val, 1)
	})

	t.Run("Next is nil if not declared", func(t *testing.T) {
		head := ll.New(1)

		AssertNil(t, head.Next)
	})
}

func TestAppend(t *testing.T) {
	testCases := []struct {
		want any
	}{
		{want: 2},
		{want: "hello"},
		{want: 'a'},
	}
	for _, tC := range testCases {
		t.Run("testing with generics", func(t *testing.T) {
			head := ll.New(1)
			head = ll.Append(head, tC.want)

			AssertEqual(t, head.Next.Val, tC.want)
		})
	}
}

func TestPrepend(t *testing.T) {
	testCases := []struct {
		want any
	}{
		{want: 2},
		{want: "hello"},
		{want: 'a'},
	}
	for _, tC := range testCases {
		t.Run("testing with generics", func(t *testing.T) {
			head := ll.New(1)
			head = ll.Prepend(head, tC.want)

			AssertEqual(t, head.Val, tC.want)
		})
	}
}

func TestReverse(t *testing.T) {
	head := ll.New(1)
	head = ll.Append(head, 2)
	head = ll.Append(head, 3)
	head = ll.Append(head, 4)

	head = ll.Reverse(head)
	AssertEqual(t, head.Val, 4)
}

func AssertEqual(t *testing.T, got, want any) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func AssertNil(t *testing.T, got *ll.ListNode) {
	t.Helper()
	if got.IsInterfaceNil() == false {
		t.Errorf("got %+v, want nil", got)
	}
}

func AssertNotEqual(t *testing.T, got, want any) {
	t.Helper()
	if got == want {
		t.Errorf("didn't want %+v", got)
	}
}
