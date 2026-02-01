package main

// Q: https://leetcode.com/problems/divide-an-array-into-subarrays-with-minimum-cost-i/?envType=daily-question&envId=2026-02-01
// Level : easy
// This is a o(N) approach

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 6, 1, 5}
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	fmt.Println(sorted)
	resultarr := []int{}
	resultarr = append(resultarr, sorted[0])
	sorted = append(sorted[:0], sorted[0+1:]...)
	sort.Ints(sorted)
	resultarr = append(resultarr, sorted[0], sorted[1])
	sum := 0
	for _, val := range resultarr {
		sum += val
	}
	fmt.Println("The result is : ", sum)
}
