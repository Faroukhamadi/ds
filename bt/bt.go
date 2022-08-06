package bt

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type TreeNode[T constraints.Ordered] struct {
	Val   T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func New[T constraints.Ordered](val T) *TreeNode[T] {
	return &TreeNode[T]{Val: val}
}

func Invert[T constraints.Ordered](root *TreeNode[T]) *TreeNode[T] {
	if root == nil {
		return nil
	}

	node := TreeNode[T]{Val: root.Val, Left: Invert(root.Right), Right: Invert(root.Left)}
	return &node
}

func Print[T constraints.Ordered](root *TreeNode[T], level int) {
	if level == 0 {
		fmt.Println("BINARY TREE")
	}
	if root != nil {
		Print(root.Right, level+1)
		prefix := ""
		for i := 0; i < level; i++ {
			prefix += "     "
		}
		fmt.Println(prefix+"->", root.Val)
		Print(root.Left, level+1)
	}
}

func IsSameTree[T constraints.Ordered](p, q *TreeNode[T]) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}

func IsSubtree[T constraints.Ordered](p, q *TreeNode[T]) bool {
	if q == nil {
		return true
	}

	if p == nil {
		return false
	}

	return IsSameTree(p, q) || IsSubtree(p.Left, q) || IsSubtree(p.Right, q)
}
