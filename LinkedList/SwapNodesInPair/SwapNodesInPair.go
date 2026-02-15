package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	// I am constructing a list here
	arr := []int{1, 2, 3, 4}
	root := &ListNode{Val: arr[0]}
	backuphead := root
	for i := 1; i < len(arr); i++ {
		root.Next = &ListNode{Val: arr[i]}
		root = root.Next
	}
	root = backuphead

	// head2 := root

	dummy := &ListNode{Val: 0, Next: root}

	prev := dummy

	// I am trying to swap the list in place
	for prev.Next != nil && prev.Next.Next != nil {
		first := prev.Next
		second := first.Next

		first.Next = second.Next
		second.Next = first

		prev.Next = second

		prev = first
	}

	// I am printing the list here
	result := dummy.Next
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}
