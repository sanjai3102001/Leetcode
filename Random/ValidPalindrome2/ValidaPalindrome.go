package main

// Q: https://leetcode.com/problems/valid-palindrome-ii/
// Level: easy
// O(N)

import "fmt"

func main() {
	s := "abc"
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			fmt.Println(ispalindrome(s, left+1, right) || ispalindrome(s, left, right-1))
		}
		left++
		right--
	}
	fmt.Println(true)

}

func ispalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
