package main

import "fmt"

func main() {
	nums := []int{0, 0, 1}
	lastnonZero := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[lastnonZero], nums[i] = nums[i], nums[lastnonZero]
			lastnonZero++
		}
	}
	fmt.Println(nums)

}
