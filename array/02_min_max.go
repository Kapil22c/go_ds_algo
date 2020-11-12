package main

import "fmt"

func minMax(a []int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

func main() {
	a := []int{5, 48, 65, 12, 35, 59}
	min, max := minMax(a)
	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)
}
