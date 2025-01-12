package main

import "fmt"

func main() {
	a := -20
	b := -30

	a = a + b
	b = a - b
	a = a - b

	fmt.Println(a)
	fmt.Println(b)
}
