package main

import "fmt"

type TreeNode struct {
	val   int
	Left  *TreeNode
	Right *TreeNode
}

func BFS(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	result := [][]int{}
	for len(queue) > 0 {
		length := len(queue)
		current := []int{}
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			current = append(current, node.val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

		}
		result = append(result, current)
	}
	return result
}

func main() {
	root := &TreeNode{val: 1}
	root.Left = &TreeNode{val: 2}
	root.Right = &TreeNode{val: 3}
	result := BFS(root)
	fmt.Println(result)

}
