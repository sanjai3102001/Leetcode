package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	arr := []int{18, 6, 10, 3}
	head := &ListNode{Val: arr[0]}
	backuphead := head
	for i := 1; i < len(arr); i++ {
		head.Next = &ListNode{Val: arr[i]}
		head = head.Next
	}
	head = backuphead
	for head != nil {
		if head.Next == nil {
			break
		}
		if head.Next != nil {
			result := gcd(head.Val, head.Next.Val)
			fmt.Println("the number is now: ", result)
			temp := &ListNode{Val: result}
			nexttemp := head.Next
			head.Next = temp
			head.Next.Next = nexttemp
		}
		head = head.Next.Next
	}
	head = backuphead
	fmt.Println("printing the resules")
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}

}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
