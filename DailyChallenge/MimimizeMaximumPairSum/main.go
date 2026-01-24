package main

import (
	"fmt"
	"sort"
)

// q: https://leetcode.com/problems/minimize-maximum-pair-sum-in-array/description/?envType=daily-question&envId=2026-01-24
// Level: Medium
// Solved in o(N)

func main() {
	nums := []int{3, 5, 2, 3}
	sort.Ints(nums)
	fmt.Println(nums)
	max := 0
	for i, j := 0, len(nums)-1; i < len(nums)/2; i, j = i+1, j-1 {
		temp := nums[i] + nums[j]
		if max < temp {
			max = temp
		}
	}
	fmt.Println(max)
}
