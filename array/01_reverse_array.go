package main

import "fmt"

func reverse(num []int) {
	for i, j := 0, len(num)-1; i < j; i, j = i+1, j-1 {
		num[i], num[j] = num[j], num[i]
	}
}

func reverse_string_datatype(fullname []string) {
	for i, j := 0, len(fullname)-1; i < j; i, j = i+1, j-1 {
		fullname[i], fullname[j] = fullname[j], fullname[i]
	}
}

func main() {
	array := []int{10, 20, 30, 40, 50}
	reverse(array)

	name := []string{"Kapil", "is", "name", "My"}

	reverse_string_datatype(name)

	fmt.Printf("%v\n", array)

	fmt.Printf("%v\n", name)
}
