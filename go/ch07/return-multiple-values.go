package main

import "fmt"

func number() (int, int) {
	return 10, 20
}

func main() {
	x, y := number()
	fmt.Println(x)
	fmt.Println(y)
}
