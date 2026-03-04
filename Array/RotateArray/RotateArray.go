package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	n := len(nums)
	k := 3
	k %= n
	temp := append(nums[n-k:], nums[:n-k]...)
	copy(nums, temp)
	fmt.Println(nums)
}
