package main

import "fmt"

func main() {
	s := "))(("
	sarr := []rune(s)
	fmt.Println(string(sarr))
	openbrace := 0
	closebrace := 0
	for i := 0; i < len(sarr); i++ {
		if sarr[i] == '(' {
			openbrace++
		} else if sarr[i] == ')' && openbrace > 0 {
			openbrace--
		} else if sarr[i] == ')' {
			closebrace++
			sarr = append(sarr[:i], sarr[i+1:]...)
			i--
		}
	}
	fmt.Println("first loop done")
	if openbrace > 0 {
		for i := len(sarr) - 1; i >= 0; i-- {
			if openbrace == 0 {
				break
			}
			if sarr[i] == '(' {
				// Remove the current '('
				sarr = append(sarr[:i], sarr[i+1:]...)
				openbrace--
			}
		}
	}
	fmt.Print("after operation: ", string(sarr))
}
