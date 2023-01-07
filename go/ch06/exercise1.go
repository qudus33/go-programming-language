package main

import "fmt"

func main() {
	number := make([]int, 3)
	fmt.Println(number)
	number[0] = 1
	number[1] = 2
	number[2] = 3
	fmt.Println(number)

	number2 := []int{1, 2, 3}
	fmt.Println(number2)
	number2 = append(number2, 4, 5)
	fmt.Println(number2)
}
