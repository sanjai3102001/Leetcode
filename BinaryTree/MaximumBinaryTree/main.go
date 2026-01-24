package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	nums := []int{3, 2, 1, 6, 0, 5}
	root := MaxBinaryTree(nums)

}

func MaxBinaryTree(nums []int) *TreeNode {
	max := 0
	maxid := 0
	for i, val := range nums {
		if max < val {
			max = val
			maxid = i
		}
	}
	fmt.Println("Max is: ", max)
	root := &TreeNode{Val: max}

	root.Left = MaxBinaryTree(nums[:maxid])
	root.Right = MaxBinaryTree(nums[maxid+1:])
	return root
}
