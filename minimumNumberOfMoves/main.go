package main

import (
	"fmt"
	"sort"
)

func main_1() {
	seats := []int{3, 1, 5}
	students := []int{2, 7, 4}
	sort.Ints(seats)
	sort.Ints(students)
	fmt.Println(seats)
	fmt.Println(students)
	result := 0
	for i := 0; i < len(seats); i++ {
		dif := seats[i] - students[i]
		if dif < 0 {
			dif = -dif
		}
		result += dif
	}
	fmt.Println(result)
}
