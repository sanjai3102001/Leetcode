package main

// Q: https://leetcode.com/problems/find-smallest-letter-greater-than-target/description/
// Level :Easy

import "fmt"

func main() {
	letters := []int{'c', 'f', 'j'}
	target := 'a'
	var result int8
	for _, val := range letters {
		if val > int(target) {
			result = int8(val)
			break
		}
	}
	fmt.Println(result)

}
