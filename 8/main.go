package main

import "fmt"

// value - установка в 1 или в 0
func setBit(num int64, i int, value int) int64 {
	if value == 1 {
		return num | (1 << i) // устанавливаем бит в 1
	} else {
		return num & ^(1 << i) // устанавливаем бит в 0
	}
}

func main() {
	var x int64
	x = 20
	fmt.Println(setBit(x, 5, 0))
}
