package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	headarr := []int{1, 2}
	// I am just makeing the liked list
	head := &ListNode{Val: headarr[0]}
	backuphead := head

	for i := 1; i < len(headarr); i++ {
		head.Next = &ListNode{Val: headarr[i]}
	}

	// I am just storing inside an array
	head = backuphead
	palindromearr := []int{}
	for head != nil {
		palindromearr = append(palindromearr, head.Val)
		head = head.Next
	}
	// Now I am just checking the palindrom using two pointer
	result := true
	for i, j := 0, len(palindromearr)-1; i < len(palindromearr)/2; i, j = i+1, j-1 {
		if palindromearr[i] != palindromearr[j] {
			result = false
			break

		}
	}
	fmt.Println("the resule is: ", result)
}
