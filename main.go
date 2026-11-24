package main

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// Boundary returns the anticlockwise boundary: root, the left edge going
// down (leaves excluded), all leaves left to right, then the right edge
// coming back up (leaves excluded). Each of the three parts is its own walk.
func Boundary(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	out := []int{root.Val}
	if root.Left == nil && root.Right == nil {
		return out // single node: it's the whole boundary, counted once
	}

	// Left edge, excluding the leaf at its bottom.
	for n := root.Left; n != nil; {
		if n.Left == nil && n.Right == nil {
			break
		}
		out = append(out, n.Val)
		if n.Left != nil {
			n = n.Left
		} else {
			n = n.Right
		}
	}

	// All leaves, left to right.
	var leaves func(n *TreeNode)
	leaves = func(n *TreeNode) {
		if n == nil {
			return
		}
		if n.Left == nil && n.Right == nil {
			out = append(out, n.Val)
			return
		}
		leaves(n.Left)
		leaves(n.Right)
	}
	leaves(root.Left)
	leaves(root.Right)

	// Right edge, collected downward then reversed, leaf excluded.
	right := []int{}
	for n := root.Right; n != nil; {
		if n.Left == nil && n.Right == nil {
			break
		}
		right = append(right, n.Val)
		if n.Right != nil {
			n = n.Right
		} else {
			n = n.Left
		}
	}
	for i := len(right) - 1; i >= 0; i-- {
		out = append(out, right[i])
	}
	return out
}

func tn(v int, l, r *TreeNode) *TreeNode { return &TreeNode{Val: v, Left: l, Right: r} }

func main() {
	root := tn(1,
		tn(2, tn(4, nil, nil), tn(5, tn(8, nil, nil), tn(9, nil, nil))),
		tn(3, tn(6, nil, nil), tn(7, nil, nil)))
	fmt.Println(Boundary(root))
}
