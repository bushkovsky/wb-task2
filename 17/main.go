package main

import (
	"fmt"
	"sort"
)

func binarySearch(sl []int, val int) int {
	sort.Ints(sl)

	first := 0
	last := len(sl) / 2

	for first <= last {
		middle := first + (last-first)/2
		if sl[middle] == val {
			return middle
		} else if sl[middle] < val {
			first = middle + 1
		} else {
			last = middle - 1
		}
	}
	return -1
}

func main() {
	slice := []int{4, 5, 1, 3, 4, 9, 1, 4, 6, 12}
	fmt.Println(binarySearch(slice, 11))
}
