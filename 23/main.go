package main

import (
	"errors"
	"fmt"
)

func delete(sl []int, ind int) ([]int, error) {
	if ind < 0 || ind >= len(sl) {
		return sl, errors.New("index out of range")
	}

	return append(sl[:ind], sl[ind+1:]...), nil

}

func main() {
	sl := []int{1, 2, 3, 4}
	sl, _ = delete(sl, 1)
	fmt.Println(sl)
}
