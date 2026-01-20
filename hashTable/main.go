package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}
	// pattern := "abb"
	temp := strings.ReplaceAll(words[2], "m", "a")
	fmt.Println(temp)

}
