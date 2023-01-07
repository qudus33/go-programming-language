package main

import "fmt"

func main() {
	numbers := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
	}
	fmt.Println(numbers["four"])
	numbers["six"] = 6
	fmt.Println(len(numbers))
	delete(numbers, "one")
	fmt.Println(len(numbers))
}
