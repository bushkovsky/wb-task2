package main

import "fmt"

func quicksort(sl []int) {
	if len(sl) <= 1 {
		return
	}

	pivot := sl[len(sl)/2]
	leftPart := make([]int, 0)
	rightPart := make([]int, 0)

	for i, value := range sl {
		if i == len(sl)/2 {
			continue
		}

		if value <= pivot {
			leftPart = append(leftPart, value)
		} else {
			rightPart = append(rightPart, value)
		}
	}
	quicksort(leftPart)
	quicksort(rightPart)
	copy(sl, append(append(leftPart, pivot), rightPart...))

}

func main() {
	sl := []int{10, 7, 8, 9, 1, 5}
	fmt.Println(sl)
	quicksort(sl)
	fmt.Println(sl)

}
