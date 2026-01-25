package main

import "fmt"

type BinaryTree struct {
	Val   int
	Left  *BinaryTree
	Right *BinaryTree
}

func main() {
	root := &BinaryTree{Val: 3}
	backuproot := root
	root.Left = &BinaryTree{Val: 2}
	root.Left.Left = &BinaryTree{Val: 1}
	root.Right = &BinaryTree{Val: 4}
	root.Right.Right = &BinaryTree{Val: 5}
	root = backuproot
	current := root

	// the real iterative starting here
	stack := []*BinaryTree{}
	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		fmt.Println(current.Val)
		current = current.Right
	}

}
