package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	// Just constructing the p
	p := &TreeNode{Val: 1}
	// backupp := p
	p.Left = &TreeNode{Val: 2}
	p.Right = &TreeNode{Val: 3}

	// Just constructing the roor2
	q := &TreeNode{Val: 1}
	// backupq := q
	q.Left = &TreeNode{Val: 2}
	q.Right = &TreeNode{Val: 3}
	result := isSameTree(p, q)
	fmt.Println(result)

}
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if q == nil || p == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
