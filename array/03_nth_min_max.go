package main

import (
	"fmt"
	"sort"
)

func nth_min(s []int, pos int) int {
	sort.Ints(s)
	return s[pos-1]
}

func main() {
	s := []int{30, 55, 1, 2, 16, 9, 27}
	pos := 3
	ans := nth_min(s, pos)
	fmt.Println(ans)
}
