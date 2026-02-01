package main

import (
	"fmt"
	"sort"
)

// root = [3,1,4,null,2], k = 1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	k := 1

	// Just creating the tree
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 4}
	// backuphead := root
	arr := helper(root)
	// Sorting the array so that we can direcly fetch the kth largest element
	sort.Ints(arr)
	fmt.Printf("the %dth Largest Node is :%d", k, arr[k-1])

}

func helper(root *TreeNode) []int {
	result := []int{}
	var traverse func(root *TreeNode)
	// Traversing and storing it inside the array
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		// post order style traversal
		traverse(root.Left)
		traverse(root.Right)
		result = append(result, root.Val)
	}
	traverse(root)
	return result
}
