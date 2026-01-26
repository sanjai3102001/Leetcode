package main

// Q: https://leetcode.com/problems/minimum-absolute-difference/submissions/1898066672/?envType=daily-question&envId=2026-01-26
// Level: easy

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{4, 2, 1, 3}
	var resultarr [][]int
	sort.Ints(arr)
	minimum := arr[1] - arr[0]
	fmt.Println(arr)
	for i, j := 0, 1; i < len(arr)-1; i, j = i+1, j+1 {
		temp := arr[j] - arr[i]
		if temp < 0 {
			fmt.Println("temp is less than 0 ", temp)
			temp = -temp
			fmt.Println("temp is less than 0 and no temp is", temp)
		}
		if temp < minimum {
			minimum = arr[j] - arr[i]
			fmt.Println("minimum is: ", minimum, "and ", arr[i], arr[j])
		}
		fmt.Println("minimum is: ", minimum, "and ", arr[i], arr[j])
	}
	fmt.Println("minimum is ", minimum)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i]-arr[j] == minimum {
				resultarr = append(resultarr, []int{arr[j], arr[i]})
			}
		}
	}
	fmt.Println(resultarr)
}
