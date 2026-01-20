package main

import "fmt"

func plusOne(digits []int) []int {
    n := len(digits)
    for i := n - 1; i >= 0; i-- {
        if digits[i] < 9 {
            digits[i]++
            return digits
        }
        digits[i] = 0
    }

    // If we are here, it means all digits were 9
    // e.g. [9,9,9] -> [1,0,0,0]
    newDigits := make([]int, n+1)
    newDigits[0] = 1
    return newDigits
}

func main() {
    // Example 1
    digits1 := []int{1, 2, 3}
    result1 := plusOne(digits1)
    fmt.Println("Example 1:", result1) // Expected: [1, 2, 4]

    // Example 2
    digits2 := []int{4, 3, 2, 1}
    result2 := plusOne(digits2)
    fmt.Println("Example 2:", result2) // Expected: [4, 3, 2, 2]

    // Example 3
    digits3 := []int{9}
    result3 := plusOne(digits3)
    fmt.Println("Example 3:", result3) // Expected: [1, 0]

    // Example 4
    digits4 := []int{9, 9, 9}
    result4 := plusOne(digits4)
    fmt.Println("Example 4:", result4) // Expected: [1, 0, 0, 0]
}