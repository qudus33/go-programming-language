package main

import "fmt"

// this program print english words of numbers from 1 to 10
func main() {
	for i := 1; i <= 12; i++ {
		switch i {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("one")
		case 2:
			fmt.Println("two")
		case 3:
			fmt.Println("three")
		case 4:
			fmt.Println("four")
		case 5:
			fmt.Println("five")
		case 6:
			fmt.Println("six")
		case 7:
			fmt.Println("seven")
		case 8:
			fmt.Println("eight")
		case 9:
			fmt.Println("nine")
		case 10:
			fmt.Println("ten")
		default:
			fmt.Println(i, "Unknown number")
		}
	}
}
