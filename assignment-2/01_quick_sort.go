package main

import (
	"fmt"
)

func partition(a []int, low, high int) int {
	p := a[high]
	for j := low; j < high; j++ {
		if a[j] < p {
			a[j], a[low] = a[low], a[j]
			low++
		}
	}

	a[low], a[high] = a[high], a[low]
	return low
}

func quickSort(a []int, low, high int) {
	if low > high {
		return
	}

	p := partition(a, low, high)
	quickSort(a, low, p-1)
	quickSort(a, p+1, high)
}

func main() {
	list := []int{18, 51, 24, 57, 89, 4, 62, 53, 23, 93, 41, 95, 84, 88}

	quickSort(list, 0, len(list)-1)
	fmt.Println(list)
}
