package bt

import (
	"fmt"
	"math"

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

func helper[T constraints.Ordered](root *TreeNode[T]) int {
	if root == nil {
		return 0
	}

	lh := helper(root.Left)
	if lh == -1 {
		return -1
	}

	rh := helper(root.Left)
	if rh == -1 {
		return -1
	}

	if math.Abs(float64(lh)-float64(rh)) > 1 {
		return -1
	}

	return int(math.Max(float64(lh), float64(rh))) + 1
}

func IsBalanced[T constraints.Ordered](root *TreeNode[T]) bool {
	h := helper(root)
	if h == -1 {
		return false
	}
	return true
}

func MaxDepth[T constraints.Ordered](root *TreeNode[T]) int {
	if root == nil {
		return 0
	}
	return max(MaxDepth(root.Left), MaxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
