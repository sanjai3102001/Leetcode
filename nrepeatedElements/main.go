package main

func main() {
	nums := []int{1, 2, 3, 3}
	length := len(nums) / 2
	for i := 0; i < len(nums); i++ {

		count := 0
		for j := 0; j < len(nums); j++ {
			if nums[i] == nums[j] {
				count++
			}
		}
		if count == length {
			println(nums[i])
			break
		}
	}

}
