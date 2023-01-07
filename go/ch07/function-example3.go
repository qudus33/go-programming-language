/*
Find the sum of the values in array.
*/
package main

import "fmt"

func sum_of_array(values []float64) {
	total := 0.0
	for i := 0; i < len(values); i++ {
		total += values[i]
	}
	fmt.Println(total)
}

func main() {
	numbers1 := []float64{10, 20, 30, 11, 21, 31}
	numbers2 := []float64{11, 21, 31, 21, 31, 41}
	sum_of_array(numbers1)
	sum_of_array(numbers2)
}
