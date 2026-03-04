package main

import "fmt"

func main() {
	nums := []int{3, 3, 6, 1}
	largest := nums[0] // the constrains says array in not empty
	for i := 1; i < len(nums); i++ {
		if largest < nums[i] {
			largest = nums[i]
		}
	}
	fmt.Println(largest)
}
