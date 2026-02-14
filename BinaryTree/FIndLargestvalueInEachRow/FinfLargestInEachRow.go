package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func FindLargest(root *TreeNode) []int {

	if root == nil {
		return []int{}
	}

	queue := []*TreeNode{root}
	result := []int{}
	for len(queue) > 0 {
		queuelength := len(queue)
		temp := queue[0]
		max := temp.Val
		for i := 0; i < queuelength; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Val > max {
				max = node.Val
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, max)
	}
	return result
}

func main() {
	root := &TreeNode{Val: 0}
	root.Left = &TreeNode{Val: -1}
	// root.Right = &TreeNode{Val: 20}
	result := FindLargest(root)
	fmt.Println(result)

}
