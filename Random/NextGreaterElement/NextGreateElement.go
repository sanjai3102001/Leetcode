package main

// Q: https://leetcode.com/problems/next-greater-element-ii/submissions/1908010020/?envType=problem-list-v2&envId=eeqswg8v
// Level: Medium
// This can be optimize

import "fmt"

func main() {
	nums := []int{5, 4, 3, 2, 1}
	resultnums := make([]int, len(nums))
	copy(resultnums, nums)
	fmt.Println("resultnums before operation: ", resultnums, "num before operation : ", nums)
	// To make the array into circular I am using index%len(nums) we can sue this method to get the array in a circular
	for i := 0; i < len(nums); i++ {
		count := 0
		var found bool
		for j := i; count < len(nums); j++ {
			if nums[j%len(nums)] > resultnums[i] {
				resultnums[i] = nums[j%len(nums)]
				found = true
				break
			}
			count++
		}
		if found != true {
			resultnums[i] = -1
		}
		fmt.Println("Num is now:", resultnums)
	}
	fmt.Println(resultnums)
}
