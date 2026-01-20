package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main_delet() {
	n := 2
	len := 0
	head := &ListNode{Val: 1}
	backup := head
	for i := 2; i < 6; i++ {
		if len == n {
			len++
			continue
		}
		len++
		head.Next = &ListNode{Val: i}
		head = head.Next
	}
	head = backup
	currentnode := head
	for currentnode != nil {
		len++
		fmt.Println("the current node is: ", currentnode.Val)
		currentnode = currentnode.Next
	}

}
