package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	left := 2
	right := 4
	// i := 0
	arr := []int{1, 2, 3, 4, 5}
	// length := len(arr)
	node := &ListNode{Val: arr[0]}
	backuphead := node
	for i := 1; i < len(arr); i++ {
		node.Next = &ListNode{Val: arr[i]}
		node = node.Next
	}
	node = backuphead
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}

	// The main code starts here
	node = backuphead
	if node == nil || left == right {
		fmt.Println("nothing to do")
	}
	dummy := &ListNode{0, node}
	prev := dummy
	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}
	current := prev.Next
	for i := 0; i < right-left; i++ {
		nextTemp := current.Next
		current.Next = nextTemp.Next
		nextTemp.Next = prev.Next
		prev.Next = nextTemp
	}

	// printing the result
	result := dummy.Next
	fmt.Println("printing the result")
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}

}
