package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
		i += 1
	}
}

// Can also be written this way;
/*
for i := 1; i <= 10; i++ {
	if i % 2 == 0 {
		fmt.Println(i, "even")
	} else {
		fmt.Println(i, "odd")
	}
}
*/
