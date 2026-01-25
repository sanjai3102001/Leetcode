package main

/*
Q:= https://leetcode.com/problems/add-two-numbers/submissions/1896681625/
Level:= Medium
*/

import (
	"fmt"
	"strconv"
)

// I am trying with the bruteforce

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// I am first adding the array to list
	l1 := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	l2 := []int{5, 6, 4}
	rootl1 := &ListNode{Val: l1[0]}
	rootl2 := &ListNode{Val: l2[0]}
	backupl1 := rootl1
	backupl2 := rootl2
	for i := 1; i < len(l1); i++ {
		rootl1.Next = &ListNode{Val: l1[i]}
		rootl1 = rootl1.Next
	}
	for i := 1; i < len(l2); i++ {
		rootl2.Next = &ListNode{Val: l2[i]}
		rootl2 = rootl2.Next
	}
	rootl1 = backupl1
	rootl2 = backupl2

	// Now I am adding the back to the array :)
	temparr1 := []int{}
	temparr2 := []int{}
	for rootl1 != nil {
		temparr1 = append(temparr1, rootl1.Val)
		rootl1 = rootl1.Next
	}
	fmt.Println(temparr1)
	for rootl2 != nil {
		temparr2 = append(temparr2, rootl2.Val)
		rootl2 = rootl2.Next
	}
	fmt.Println(temparr2)

	// concerting to string
	int1 := 0
	int2 := 0
	for i := len(temparr1) - 1; i >= 0; i-- {
		num := temparr1[i]
		int1 = int1*10 + num
	}
	fmt.Println(int1)

	for i := len(temparr2) - 1; i >= 0; i-- {
		num := temparr2[i]
		int2 = int2*10 + num
	}
	fmt.Println(int2)
	sum := int1 + int2
	fmt.Println("sum is : ", sum)
	strsum := strconv.Itoa(sum)
	temp, _ := strconv.Atoi(string(strsum[len(strsum)-1]))
	resultlist := &ListNode{Val: temp}
	backuphead := resultlist
	for i := len(strsum) - 2; i >= 0; i-- {
		temp, _ = strconv.Atoi(string(strsum[i]))
		resultlist.Next = &ListNode{Val: temp}
		resultlist = resultlist.Next
	}
	resultlist = backuphead

	// printing the result list
	fmt.Println("Printing the list")
	for resultlist != nil {
		fmt.Println(resultlist.Val)
		resultlist = resultlist.Next
	}
}
