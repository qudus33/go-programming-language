package main

import "fmt"

func main() {
	fmt.Println("This program multiply any unmber by 2")
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2
	fmt.Println(output)
}
