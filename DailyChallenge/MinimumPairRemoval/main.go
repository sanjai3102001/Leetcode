package main

import "fmt"

func main() {
	nums := []int{5, 2, 3, 1}
	count := 0
	assending := true
	sum := 0
outer:
	for {
		fmt.Println(nums)
	checkassending:
		for i := 0; i < len(nums); i++ {
			prev := nums[0]
			if len(nums) < 2 {
				break outer
			} else if prev < nums[i] {
				assending = false

			}
		}
		if assending == false {
			goto checkassending
		}
		sum = nums[0] + nums[1]
	starteover:
		for i := 1; i < len(nums)-1; i++ {
			temp := nums[i] + nums[i+1]
			if sum < temp {
				count++
				sum = temp
				nums = append(nums[:i], nums[i+2:]...)
				goto starteover
			}

		}
	}
	fmt.Println("Count is: ", count)
}
