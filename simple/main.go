package main

import (
	"fmt"
	"strings"
)

func main_old() {
	s := "leet**cod*e"
	var stackarr []string

	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			stackarr = stackarr[:len(stackarr)-1]
		} else {
			stackarr = append(stackarr, string(s[i]))
		}
	}
	result := strings.Join(stackarr, "")
	fmt.Println((stackarr))
	fmt.Println(result)
	// // s = strings.ReplaceAll(s, s[4-1:4], "")
	// i := 0
	// for 1 == 1 {
	// 	if s[i] == '*' {
	// 		s = strings.ReplaceAll(s, s[i-1:i+1], "")
	// 		i = 0
	// 	}
	// 	if i == len(s)-1 {
	// 		break
	// 	}
	// 	i++
	// }
	// fmt.Println(s)
	// // i := 0
	// // for 1 == 1 {
	// // 	if s[i] == '*' {
	// // 		s = s[:i-1] + s[i+1:]
	// // 		fmt.Println(s)
	// // 		i = 0
	// // 	}
	// // 	if i == len(s)-1 {
	// // 		break
	// // 	}
	// // 	i++
	// // }

}
