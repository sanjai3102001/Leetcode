package main

import "fmt"

type List struct {
	value string
	next  *List
}
type List2 struct {
	value string
	next  *List
}

func mainSample() {
	node1 := &List{value: "first"}
	node2 := &List{value: "second"}
	node3 := &List{value: "Third"}
	target := 2
	len := 1
	node1.next = node2
	node2.next = node3
	current := node1
	listnode2 := node1
	for current != nil {
		if len == target {
			listnode2 = current.next
		}
		len++
		fmt.Println("the node 2 is now: ", listnode2.value)
		current = current.next

	}
}
