package main

// Q: https://leetcode.com/problems/reverse-only-letters/
// Level  Easy

import (
	"fmt"
	"unicode"
)

func main() {
	s := "ab-cd"
	sarr := []rune(s)
	for i, j := 0, len(sarr)-1; i < j; {
		if !unicode.IsLetter(sarr[i]) {
			i++
			continue
		} else if !unicode.IsLetter(sarr[j]) {
			j--
			continue
		}
		sarr[i], sarr[j] = sarr[j], sarr[i]
		i++
		j--
	}
	result := string(sarr)
	fmt.Println(result)
}
