package main

import "fmt"

func main() {
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 1))

	a := 0
	increment := func() int {
		a++
		return a
	}
	fmt.Println(increment())
	fmt.Println(increment())
	// A function like this together with the non-local variables it references is known as a closure. In this case increment and the variable a form the closure.
}
