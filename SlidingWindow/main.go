package main

import "fmt"

func main_2() {
	nums := []int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}
	k := 2
	right := k
	left := 0
	nice := 0

	for 1 == 1 {
		count := 0
		for i := left; i < right; i++ {
			if nums[i]%2 != 0 {
				count++
			}
		}
		if count == k {
			left++
			nice++
		} else if count < k && left > 0 {
			right++
		} else if count > k {
			left++
		} else if count < k {
			right++
		}
		if right >= len(nums) {
			break

		}
	}
	fmt.Println("the nice pair is: ", nice)

}
