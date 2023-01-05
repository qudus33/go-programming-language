package main

import (
	"fmt"
)

func main() {
	var x string = "Hello World"
	fmt.Println(x)

	var name string
	name = "Abdulqudus"
	fmt.Println(name)

	var age = "ten"
	fmt.Println(age)

	var first = "First "
	first += "Second"
	fmt.Println(first)

	var i string = "hello"
	var j string = "world"
	fmt.Println(i == j)

	var k = "hello"
	var l = "hello"
	fmt.Println(k == l)

	username := "qudus33"
	fmt.Println(username)

	lenght := 10
	breadth := 5
	area := lenght * breadth
	fmt.Println("area =", area)
}
