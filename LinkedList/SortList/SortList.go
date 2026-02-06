package main

// Q:=https://leetcode.com/problems/sort-list/description/
// Level: Medium

import (
	"fmt"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	arr := []int{4, 2, 1, 3}

	// makeing the list
	head := &ListNode{Val: arr[0]}
	backuphead := head
	for _, i := range arr {
		head.Next = &ListNode{Val: i}
		head = head.Next
	}

	// The solution starts here
	sortarr := []int{}
	head = backuphead
	// if head == nil {
	// 	fmt.Println("head is nil so and wil be nil")
	// }
	for head != nil {
		sortarr = append(sortarr, head.Val)
		head = head.Next
	}

	sort.Ints(sortarr)
	fmt.Println(sortarr)

	// Adding to the list
	resulthead := &ListNode{Val: sortarr[0]}
	backupresulthead := resulthead
	for i := 1; i < len(sortarr); i++ {
		resulthead.Next = &ListNode{Val: sortarr[i]}
		resulthead = resulthead.Next
	}
	resulthead = backupresulthead

	// Printing the linkedlist
	for resulthead != nil {
		fmt.Println(resulthead.Val)
		resulthead = resulthead.Next
	}

}
