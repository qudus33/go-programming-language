package main

import "fmt"

func get_name(name string) {
	fmt.Println("My name is", name)
}

func main() {
	get_name("Abdul")
	get_name("Qudus")
	get_name("Abdul Qudus")
}
