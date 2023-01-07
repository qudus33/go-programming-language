package main

import "fmt"

func main() {
	x := []int{
		48, 96, 86, 68, 57,
		82, 63, 70, 3, 34, 83,
		27, 19, 97, 9, 17, 1, 2,
	}
	fmt.Println(x)
	smallest := x[0]
	for i := 0; i < len(x); i++ {
		if x[i] < smallest {
			smallest = x[i]
		}
	}
	fmt.Println("The smallest number is:", smallest)

}
