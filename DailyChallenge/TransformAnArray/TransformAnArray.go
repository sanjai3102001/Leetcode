package main

// Q:=https://leetcode.com/problems/transformed-array/?envType=daily-question&envId=2026-02-05
// Level: Easy

import "fmt"

// formula backward ((i-k)%n+n)%n > if the index is in negative useing this ((i+k)%n+n)%n > we know the the number will be less than 0 so useing the second one
// formula forward (i+k)%n

func main() {
	nums := []int{-1, 4, -1}
	result := make([]int, len(nums))
	length := len(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			result[i] = nums[(i+nums[i])%length]
		} else if nums[i] < 0 {
			result[i] = nums[((i+nums[i])%length+length)%length]
		} else {
			result[i] = nums[i]
		}
		fmt.Println("The result is now: ", result)
	}
	fmt.Println(result)
}
