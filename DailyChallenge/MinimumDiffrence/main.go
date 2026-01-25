package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// nums := []int{87063, 61094, 44530, 21297, 95857, 93551, 9918}
	nums := []int{90}
	if len(nums) <= 1 {
		fmt.Println("the length is 0")
	}
	sort.Ints(nums)
	fmt.Println(nums)

	minDif := math.MaxInt32
	k := 1
	for i := 0; i <= len(nums)-k; i++ {
		start := i
		end := i + k - 1
		diff := nums[end] - nums[start]
		if diff <= minDif {
			minDif = diff
		}

	}
	fmt.Println(minDif)
}
