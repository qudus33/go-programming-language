package main

import "fmt"

func main() {
	var x = make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])
}
