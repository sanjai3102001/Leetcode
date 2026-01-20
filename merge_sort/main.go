package main

import "fmt"

func main() {
	arr := []int{1, 3, 4, 6, 2, 4, 5, 2, 9, 2}
	fmt.Println("before merge: ", arr)
	arr = mergesort(arr)
	fmt.Println("after sort: ", arr)
}

func mergesort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	firstHalf := arr[:mid]
	secondHalf := arr[mid:]
	sortedLeft := mergesort(firstHalf)
	sortedRight := mergesort(secondHalf)
	return merge(sortedLeft, sortedRight)

}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}
