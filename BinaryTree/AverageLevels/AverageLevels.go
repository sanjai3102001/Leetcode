package main

import "fmt"

// import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func AverageInLevel(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}
	queue := []*TreeNode{root}
	result := []float64{}

	for len(queue) > 0 {
		queuelength := len(queue)
		current := []int{}
		var sum float64
		var count float64
		for i := 0; i < queuelength; i++ {
			node := queue[0]
			queue = queue[1:]
			current = append(current, node.Val)
			sum += float64(node.Val)
			count++

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		fmt.Println(float64(sum / count))
		result = append(result, float64(sum/count))
	}
	return result
}

func main() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}

	result := AverageInLevel(root)
	// print(result)
	for _, val := range result {
		fmt.Printf("%f", val)
	}
}
