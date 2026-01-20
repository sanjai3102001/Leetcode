package main

import "fmt"

func main() {
	s := "abcabc"
	sarr := []rune(s)
	letter := []int{0, 0, 0}
	count := 0
	right := 0
	left := 0
	fmt.Println(string(sarr))
	for ; right < len(sarr); right++ {
		switch sarr[right] {
		case 'a':
			letter[0]++
		case 'b':
			letter[1]++
		case 'c':
			letter[2]++
		}
		for letter[0] > 0 && letter[1] > 0 && letter[2] > 0 {
			count += len(sarr) - right
			switch sarr[left] {
			case 'a':
				letter[0]--
			case 'b':
				letter[1]--
			case 'c':
				letter[2]--

			}
			left++
		}
	}
	fmt.Println("The count is: ", count)
}
