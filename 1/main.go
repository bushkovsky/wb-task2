package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human
(аналог наследования).
*/

// Human структура, которую будем встраивать
type Human struct {
	age  uint
	name string
}

// Greeting - метод приветствия
func (h Human) Greeting() {
	fmt.Println("Hello, I'm " + h.name)
}

// Action - куда по заданию надо встроить Human
type Action struct {
	Human
}

func main() {
	a := new(Action)
	a.age = 10
	a.name = "Active"
	a.Greeting() // для экземпляра Action доступен метод Greeting структуры Human
}
