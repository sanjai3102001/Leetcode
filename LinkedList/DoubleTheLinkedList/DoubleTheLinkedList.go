package main

// Q:= https://leetcode.com/problems/double-a-number-represented-as-a-linked-list/description/
// Level: medium

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// I am constructing a linkedlist
	arr := []int{1, 8, 9}
	head := &ListNode{Val: arr[0]}
	backuphead := head
	for i := 1; i < len(arr); i++ {
		head.Next = &ListNode{Val: arr[i]}
		head = head.Next
	}
	head = backuphead

	if head.Val > 4 {
		newhead := &ListNode{Val: 0, Next: head}
		head = newhead
	}

	current := head

	for current != nil {

		current.Val = (current.Val * 2) % 10
		if current.Next != nil && current.Next.Val > 4 {
			current.Val += 1
		}
		current = current.Next
	}

	// Now just prinitng the list
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
