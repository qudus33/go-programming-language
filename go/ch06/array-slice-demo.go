package main

import "fmt"

func main() {
	x := [4]string{"ball", "cat", "kettle", "fan"}
	fmt.Println(x)
	for i := 0; i < len(x); i++ {
		note := "This is a "
		fmt.Println(note + x[i])
	}

	y := [5]int{1, 2, 3, 4, 5}
	fmt.Println(y)
	sum := 0
	for i := 0; i < len(y); i++ {
		sum += y[i]
	}
	fmt.Println("The sum of values in y is:", sum)

	// Create an array with the make function
	z := make([]string, 2)
	z[0] = "tea"
	z[1] = "pot"
	fmt.Println(z)
	fmt.Println("This is a " + z[0] + z[1])

	var d [5]string
	d[0] = "This"
	d[1] = "is"
	d[2] = "a"
	d[3] = "mango"
	d[4] = "tree"
	fmt.Println(d)
	// slicing some value from d and append it to a value.
	e := d[2:5]
	fmt.Println(e)
	e = append(e, "in", "the", "bush")
	fmt.Println(e)

	// using the copy function
	g := make([]string, 2)
	copy(g, e[4:6])
	fmt.Println(g)
}
