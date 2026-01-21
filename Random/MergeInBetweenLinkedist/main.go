package main

import "fmt"

//Q: https://leetcode.com/problems/merge-in-between-linked-lists/description/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	arr1 := []int{10, 1, 13, 6, 9, 5}
	list1 := &ListNode{Val: arr1[0]}
	backuphead1 := list1
	fmt.Println("Starting....")
	for i := 1; i < len(arr1); i++ {
		list1.Next = &ListNode{Val: arr1[i]}
		list1 = list1.Next
	}
	list1 = backuphead1.Next
	arr2 := []int{1000000, 1000001, 1000002}
	list2 := &ListNode{Val: arr2[0]}
	backuphead2 := list2
	for i := 1; i < len(arr2); i++ {
		list2.Next = &ListNode{Val: arr2[i]}
		list2 = list2.Next
	}
	list2 = backuphead2

	// Tryinng with bruteforce approch for now. this needs to br optimized and try to do inplace to optimize memory
	a := 3
	b := 4
	count := 1
	ResultList := &ListNode{Val: arr1[0]}
	backuprsult := ResultList
	for ResultList != nil {
		if count < a {
			ResultList.Next = list1
			ResultList = ResultList.Next
			list1 = list1.Next
			count++
			continue
		} else if count >= a && count <= b {
			count++
			list1 = list1.Next
			fmt.Println("I am running in a loop and count is now: ", count)
			continue
		} else if count > b && list2 != nil {
			ResultList.Next = list2
			ResultList = ResultList.Next
			list2 = list2.Next
			count++
			continue
		} else if list1 != nil {
			ResultList.Next = list1
			ResultList = ResultList.Next
		}
		break
	}
	ResultList = backuprsult

	// I am now just printinhg the list
	fmt.Println("Printing the final result")
	ResultList = backuprsult
	for ResultList != nil {
		fmt.Println(ResultList.Val)
		ResultList = ResultList.Next
	}
}
