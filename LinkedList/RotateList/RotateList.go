package main

// Q:= https://leetcode.com/problems/rotate-list/
// Level:= Medium
// This is a kind of O(1)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	// making the list
	arr := []int{0, 1, 2}
	k := 4
	head := &ListNode{Val: arr[0]}
	backuphead := head
	for i := 1; i < len(arr); i++ {
		head.Next = &ListNode{Val: arr[i]}
		head = head.Next
	}
	head = backuphead
	if head == nil {
		fmt.Println("head is nil")
	}
	// Placing back to array
	resultarr := []int{}
	for head != nil {
		resultarr = append(resultarr, head.Val)
		head = head.Next
	}

	// Performing the real calculation in array
	n := len(arr)
	k = k % n
	resultarr = append(resultarr[n-k:], resultarr[:n-k]...)
	fmt.Println(resultarr)

	// Placing in an list
	head = backuphead
	head = &ListNode{Val: resultarr[0]}
	backuphead = head
	for i := 1; i < len(resultarr); i++ {
		head.Next = &ListNode{Val: resultarr[i]}
		head = head.Next
	}

	// printing the list just for my refrence
	head = backuphead
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
