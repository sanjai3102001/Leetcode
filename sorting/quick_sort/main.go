package main

import "fmt"

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	fmt.Println("original nums: ", nums)
	quicksort(nums, 0, len(nums)-1)
	fmt.Print("after sort: ", nums)
}

func quicksort(nums []int, low, high int) {
	if low < high {
		pivotIndex := partition(nums, low, high)
		quicksort(nums, low, pivotIndex-1)
		quicksort(nums, pivotIndex+1, high)
	}
}

func partition(nums []int, low, high int) int {
	pivot := nums[high]
	i := low - 1

	for j := low; j < high; j++ {
		if nums[j] < pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[high] = nums[high], nums[i+1]
	return i + 1

}
