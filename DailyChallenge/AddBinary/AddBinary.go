package main

import (
	"fmt"
	"slices"
)

func main() {
	a := "11"
	b := "1"
	carry := 0
	result := []byte{}
	i, j := len(a)-1, len(b)-1
	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}

		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}

		result = append(result, byte(sum%2+'0'))
		carry = sum / 2

	}
	slices.Reverse(result)
	fmt.Println(string(result))

}
