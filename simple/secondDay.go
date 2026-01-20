package main

import (
	"fmt"
)

func main() {
	s := "abab"
	mymap := make(map[string]struct{})
	for i := 0; i < len(s); i++ {
		mymap[string(s[i])] = struct{}{}

	}
	fmt.Println(len(mymap))
}
