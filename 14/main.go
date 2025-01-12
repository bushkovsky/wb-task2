package main

import "fmt"

/*
Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}.
*/

func typeCheck(value any) {

	switch value.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("chan int")
	case chan bool:
		fmt.Println("chan bool")
	default:
		fmt.Println("other")

	}
}

func main() {
	typeCheck(5)
	typeCheck("uu")
	typeCheck(make(chan int))
}
