package main

import "fmt"

// If you are seeing this this is not solved with two pointer try to solve using two pointer
// tat case "][[]][][[][]"
func main1() {
	s := "]]][[["
	sarr := []rune(s)
	unmatchedopen := 0
	unmatchclose := 0
	for i := 0; i < len(sarr); i++ {
		if sarr[i] == '[' {
			unmatchedopen++
		} else {
			if unmatchedopen > 0 {
				unmatchedopen--
			} else {
				unmatchclose++
			}
		}
	}
	fmt.Println("mached oper: ", unmatchedopen, "and matched close: ", unmatchclose)
	fmt.Println((unmatchedopen + 1) / 2)
}
