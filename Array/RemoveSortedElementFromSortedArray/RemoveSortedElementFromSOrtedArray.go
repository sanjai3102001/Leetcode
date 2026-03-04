package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	i := 1
	for i < len(nums) {
		if nums[i-1] == nums[i] {
			nums = append(nums[:i-1], nums[i:]...)
			i--
		}
		i++
	}
	fmt.Println(nums)
}
