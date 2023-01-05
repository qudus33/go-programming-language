package main

import "fmt"

func main() {
	fmt.Println("This program converts from feets into meters.")
	fmt.Print("Enter the value in feet: ")
	var feet float64
	fmt.Scanf("%f", &feet)

	meters := (feet * 0.3048)
	fmt.Println("The value in meters is", meters)
}
