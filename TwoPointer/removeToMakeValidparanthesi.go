package main

import "fmt"

func main() {
	s := "lee(t(c)o)de)"
	sarr := []rune(s)
	notmatchingopen := 0
	// matchingopen := 0

	// notmatchingcllose := 0
	for i := 0; i < len(sarr); i++ {
		if sarr[i] == '[' {
			notmatchingopen++
		} else {
			if sarr[i] == ']' && notmatchingopen > 0 {
				notmatchingopen--
			} else {
				if sarr[i] == ']' {
					sarr = append(sarr[:i-1], sarr[i:]...)
				}

			}
		}
	}
	fmt.Println(string(sarr))

}
