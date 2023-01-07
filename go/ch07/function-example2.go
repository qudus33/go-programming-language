package main

import (
	"fmt"
)

func sum(x, y int) int {
	add_number := x + y
	return add_number
}

func main() {
	fmt.Println(sum(2, 3))
	fmt.Println(sum(10, 12))
	fmt.Println(sum(14, 34))
}
