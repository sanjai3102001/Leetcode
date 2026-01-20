package main

import "fmt"

type LinkedList struct {
	Val  int
	Next *LinkedList
}

func main_sum() {
	arr := []int{0, 3, 1, 0, 4, 5, 2, 0}
	node := &LinkedList{Val: arr[0]}
	backuphead := node
	for i := 1; i < len(arr); i++ {
		node.Next = &LinkedList{Val: arr[i]}
		node = node.Next
	}
	node = backuphead
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}

	// i := 0
	sum := 0
	isend := false
	sumarr := []int{}
	node = backuphead
	var resulelist *LinkedList
	var resulthead *LinkedList
	// the main task comes here
	for node != nil {

		if node.Val == 0 && isend == true {
			sumarr = append(sumarr, sum)
			if resulelist == nil {
				resulelist = &LinkedList{Val: sum}
				resulthead = resulelist
			} else {
				resulelist.Next = &LinkedList{Val: sum}
				resulelist = resulelist.Next
			}
			sum = 0
			// node = node.Next
			isend = false
		} else if node.Val == 0 && isend == false {
			isend = true
		}
		sum += node.Val
		fmt.Println("the node is: ", node.Val)
		node = node.Next

	}
	resulelist = resulthead
	fmt.Println(sumarr)

}
