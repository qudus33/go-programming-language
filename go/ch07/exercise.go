package main

import "fmt"

// 4) Using makeEvenGenerator as an example, write a makeOddGenerator function that generates odd numbers.

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret += i
		i += 2
		return
	}
}
func main() {
	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd()) //1
	fmt.Println(nextOdd()) //3
	fmt.Println(nextOdd()) //5
}

// 3) Write a function with one variadic parameter that finds the greatest number in a list of numbers.
/*
func greates_number(args ...int) int {
	number := args[0]
	for _, v := range args {
		if v > number {
			number = v
		}
	}
	return number
}

func main() {
	fmt.Println(greates_number(11, 47, 43, 5, 64, 32, 56))
}
*/

/*
// 2) Write a function which takes an integer and halves it and returns true if it was even or false if it was odd. For example half(1) should return (0, false) and half(2) should return (1, true).

func half(number int) bool {
	get_half := number % 2
	if get_half == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(half(2))
}
*/

/*
// 1) sum is a function which takes a slice of numbers and adds them together. What would its function signature look like in Go?

func sum(numbers []int) int {
	total := 0
	for i := 0; i < len(numbers); i++ {
		total = total + numbers[i]
	}
	return total
}

func main() {
	fmt.Println(sum([]int{1, 2, 3, 4, 5, 6}))
}
*/
