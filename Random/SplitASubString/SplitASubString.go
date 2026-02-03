package main

// Q: https://leetcode.com/problems/split-a-string-in-balanced-strings/submissions/1907140130/
// Level: Easy

func main() {
	s := "RLRRLLRLRL"
	ans := 0
	var i int
	for _, val := range s {
		if val == 'L' {
			i++
		} else {
			i--
		}
		if i == 0 {
			ans++
		}
	}
	println(ans)
}
