package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func RightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	count := 0
	queue := []*TreeNode{root}
	result := []int{}
	for len(queue) > 0 {
		length := len(queue)
		current := []int{}
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			current = append(current, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

		}
		temp := current[len(current)-1]
		result = append(result, temp)
		count++

	}
	return result
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 4}
	result := RightSideView(root)
	fmt.Println(result)

	final := make([]int, 0)
	for _, val := range result {
		final = append(final, val)
	}
	fmt.Println("the final output is ")
	fmt.Println(final)

}
