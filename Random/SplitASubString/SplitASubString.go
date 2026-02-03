package main

// Q: https://leetcode.com/problems/split-a-string-in-balanced-strings/submissions/1907140130/
// Level: Easy

func main() {
	s := "RLRRLLRLRL"
	ans := 0
	var i int
	for _, val := range s {
		// Trackk pair Add L and Remove R
		if val == 'L' {
			i++
		} else {
			i--
		}
		// If ) we found a pair add(1) to answer
		if i == 0 {
			ans++
		}
	}
	println(ans)
}
