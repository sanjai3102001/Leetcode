package main

// Q: https://leetcode.com/problems/remove-nodes-from-linked-list/submissions/1893707348/?envType=problem-list-v2&envId=linked-list

// This is a brute force approach , need to optimize more I have less time today need to optimize more

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	arr := []int{5, 2, 13, 3, 8}
	head := &ListNode{Val: arr[0]}
	backuphead1 := head
	fmt.Println("Starting....")
	for i := 1; i < len(arr); i++ {
		head.Next = &ListNode{Val: arr[i]}
		head = head.Next
	}
	head = backuphead1

	// actual ans
	// result := &ListNode{}
	resultarr := []int{}
	i := 0
	for head != nil {
		resultarr = append(resultarr, head.Val)
		head = head.Next
		i++
	}
	fmt.Println(resultarr)

	needtoreverse := []int{}
	// needtoreverse = append(needtoreverse, resultarr[len(resultarr)-1])
	max := 0
	for i := len(resultarr) - 1; i >= 0; i-- {

		if resultarr[i] >= max {
			max = resultarr[i]
			needtoreverse = append(needtoreverse, max)
		}
	}
	fmt.Print(needtoreverse)

	resultNode := &ListNode{Val: needtoreverse[len(needtoreverse)-1]}
	backupresultnode := resultNode
	for i := len(needtoreverse) - 2; i >= 0; i-- {
		resultNode.Next = &ListNode{Val: needtoreverse[i]}
		resultNode = resultNode.Next
	}

	resultNode = backupresultnode
	// I am printing the linked list
	for resultNode != nil {
		fmt.Println(resultNode.Val)
		resultNode = resultNode.Next
	}
}
