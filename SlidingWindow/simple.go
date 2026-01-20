package main

import "fmt"

func main_changing() {
	arr := []int{1, 2, 3, 4, 5}
	k := 3
	currentsum := 0
	for i := 0; i < k; i++ {
		currentsum += arr[i]
	}
	maxsum := currentsum
	for i := k; i < len(arr); i++ {
		currentsum += arr[i] - arr[i-k]
		if currentsum > maxsum {
			maxsum = currentsum
		}
	}
	fmt.Println("max sum is: ", maxsum)
}
