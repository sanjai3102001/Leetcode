package main

import (
	"fmt"
)

func main() {
	widths := []int{3, 4, 10, 4, 8, 7, 3, 3, 4, 9, 8, 2, 9, 6, 2, 8, 4, 9, 9, 10, 2, 4, 9, 10, 8, 2}
	s := "mqblbtpvicqhbrejb"
	fmt.Println(102 % 100)
	result := 0

	// I am just making map
	mymap := make(map[rune]int)
	temp := 0
	fmt.Println("printing my map")
	for i := 97; i < 123; i++ {
		mymap[rune(i)] = widths[temp]
		temp++
	}
	fmt.Println(mymap)
	temp = 0
	//calculating total value
	resultarr := make([]int, 2)
	for _, j := range s {
		temp += mymap[rune(j)]
		if temp > 100 {
			resultarr[0] += 1
			temp = mymap[rune(j)]
		} else if temp == 100 {
			resultarr[0] += 1
			temp = 0
		}
		result += mymap[rune(j)]

	}
	resultarr[0] += 1
	resultarr[1] = temp

	// fmt.Println("Total number is :", result)
	// fmt.Println(result)
	// if result%100 != 0 {
	// 	resultarr[0] = (result / 100) + 1
	// 	resultarr[1] = result % 100
	// } else if result%100 == 0 {
	// 	resultarr[0] = result / 100
	// 	resultarr[1] = 0
	// }
	fmt.Println(resultarr)
}
