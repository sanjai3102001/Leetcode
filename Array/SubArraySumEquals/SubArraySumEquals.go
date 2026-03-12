package main

import "fmt"

func main() {
	nums := []int{1, 1, 1}
	k := 2
	mymap := make(map[int]int)
	mymap[0] = 1
	count := 0
	currentsum := 0
	for i := 0; i < len(nums); i++ {
		currentsum += nums[i]
		freq, ok := mymap[currentsum-k]
		if ok {
			count += freq
		}
		mymap[currentsum]++
	}
	fmt.Println(count)
}
