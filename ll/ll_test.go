package ll_test

import (
	"github.com/Faroukhamadi/ds/ll"
	"testing"
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
